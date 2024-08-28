/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package oracle

import (
	"bytes"
	"crypto/rand"
	"crypto/x509"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gravitational/trace"
	"github.com/pavlo-v-chernykh/keystore-go/v4"
	"software.sslmate.com/src/go-pkcs12"

	"github.com/gravitational/teleport"
	"github.com/gravitational/teleport/api/constants"
	"github.com/gravitational/teleport/api/identityfile"
	"github.com/gravitational/teleport/api/utils/keys"
	"github.com/gravitational/teleport/lib/client"
	"github.com/gravitational/teleport/lib/tlsca"
	"github.com/gravitational/teleport/lib/utils"

	clientidentityfile "github.com/gravitational/teleport/lib/client/identityfile"
)

// GenerateClientConfiguration function generates following Oracle Client configuration:
// wallet.jks   - Java Wallet format used by JDBC Drivers.
// sqlnet.ora   - Generic Oracle Client Configuration File allowing to specify Wallet Location.
// tnsnames.ora - Oracle Net Service mapped to connections descriptors.
func GenerateClientConfiguration(key *client.Key, db tlsca.RouteToDatabase, profile *client.ProfileStatus) error {
	walletPath := profile.OracleWalletDir(key.ClusterName, db.ServiceName)
	if err := os.MkdirAll(walletPath, teleport.PrivateDirMode); err != nil {
		return trace.Wrap(err)
	}
	password, err := utils.CryptoRandomHex(32)
	if err != nil {
		return trace.Wrap(err)
	}

	localProxyCAPem, err := os.ReadFile(profile.DatabaseLocalCAPath())
	if err != nil {
		return trace.ConvertSystemError(err)
	}

	jksWalletPath, err := createClientWallet(key, localProxyCAPem, password, walletPath)
	if err != nil {
		return trace.Wrap(err)
	}

	if runtime.GOOS == constants.WindowsOS {
		jksWalletPath = strings.ReplaceAll(jksWalletPath, `\`, `\\`)
	}

	err = writeClientConfig(walletPath, jksWalletPath, password)
	if err != nil {
		return trace.Wrap(err)
	}
	return nil
}

func createClientWallet(key *client.Key, certPem []byte, password string, walletPath string) (string, error) {
	buff, err := createJKSWallet(key.PrivateKeyPEM(), certPem, certPem, password)
	if err != nil {
		return "", trace.Wrap(err)
	}
	walletFile := filepath.Join(walletPath, "wallet.jks")
	if err := os.WriteFile(walletFile, buff, teleport.FileMaskOwnerOnly); err != nil {
		return "", trace.Wrap(err)
	}

	if clientidentityfile.IsOrapkiAvailable() {
		if _, err := createOracleWallet(walletPath, key.PrivateKeyPEM(), certPem, certPem, password); err != nil {
			return "", trace.Wrap(err)
		}
	}
	return walletFile, nil
}
func createOracleWallet(path string, keyPEM, certPEM, caPEM []byte, password string) ([]byte, error) {
	certBlock, err := tlsca.ParseCertificatePEM(certPEM)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	keyK, err := keys.ParsePrivateKey(keyPEM)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	var caCerts []*x509.Certificate
	c, err := tlsca.ParseCertificatePEM(caPEM)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	caCerts = append(caCerts, c)

	pf, err := pkcs12.LegacyRC2.WithRand(rand.Reader).Encode(keyK.Signer, certBlock, nil, password)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	p12Path := filepath.Join(path, "ca.p12")
	certPath := filepath.Join(path, "user.crt")

	if err := os.WriteFile(p12Path, pf, identityfile.FilePermissions); err != nil {
		return nil, trace.Wrap(err)
	}
	err = os.WriteFile(certPath, caPEM, identityfile.FilePermissions)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if err := os.MkdirAll(path, teleport.PrivateDirMode); err != nil {
		return nil, trace.Wrap(err)
	}

	if err := clientidentityfile.CreateOracleWallet([]string{}, path, p12Path, password); err != nil {
		return nil, trace.Wrap(err)
	}
	return nil, nil
}

func createJKSWallet(keyPEM, certPEM, caPEM []byte, password string) ([]byte, error) {
	key, err := keys.ParsePrivateKey(keyPEM)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	privateKey, err := x509.MarshalPKCS8PrivateKey(key.Signer)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	ks := keystore.New()
	pkeIn := keystore.PrivateKeyEntry{
		CreationTime: time.Now(),
		PrivateKey:   privateKey,
		CertificateChain: []keystore.Certificate{
			{
				Type:    "x509",
				Content: certPEM,
			},
		},
	}

	if err := ks.SetPrivateKeyEntry("teleportUserCert", pkeIn, []byte(password)); err != nil {
		return nil, trace.Wrap(err)
	}
	trustIn := keystore.TrustedCertificateEntry{
		CreationTime: time.Now(),
		Certificate: keystore.Certificate{
			Type:    "x509",
			Content: caPEM,
		},
	}
	if err := ks.SetTrustedCertificateEntry("teleportLocalCA", trustIn); err != nil {
		return nil, trace.Wrap(err)
	}
	var buff bytes.Buffer
	if err := ks.Store(&buff, []byte(password)); err != nil {
		return nil, trace.Wrap(err)
	}
	return buff.Bytes(), nil
}

func writeClientConfig(path string, jksFile string, password string) error {
	var clientConfiguration = []templateSettings{
		tnsNamesORASettings{
			Host: "localhost",
			// User default values that will be overwritten by JDBC connection string.
			ServiceName: "XE",
			Port:        "2484",
		},
		sqlnetORASettings{
			WalletDir: path,
		},
		jdbcSettings{
			KeyStoreFile:       jksFile,
			TrustStoreFile:     jksFile,
			KeyStorePassword:   password,
			TrustStorePassword: password,
		},
	}

	for _, v := range clientConfiguration {
		if err := writeSettings(v, path); err != nil {
			return trace.Wrap(err, "Failed to write %v", v.configFilename())
		}
	}
	return nil
}
