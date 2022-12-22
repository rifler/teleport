// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authz_test

import (
	"testing"

	"github.com/gravitational/trace"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh"

	"github.com/gravitational/teleport"
	"github.com/gravitational/teleport/api/constants"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/devicetrust/authz"
	"github.com/gravitational/teleport/lib/modules"
	"github.com/gravitational/teleport/lib/tlsca"
)

func TestVerifyTLSUser(t *testing.T) {
	runVerifyUserTest(t, "VerifyTLSUser", func(dt *types.DeviceTrust, ext *tlsca.DeviceExtensions) error {
		return authz.VerifyTLSUser(dt, tlsca.Identity{
			Username:         "llama",
			DeviceExtensions: *ext,
		})
	})
}

func TestVerifySSHUser(t *testing.T) {
	runVerifyUserTest(t, "VerifySSHUser", func(dt *types.DeviceTrust, ext *tlsca.DeviceExtensions) error {
		return authz.VerifySSHUser(dt, &ssh.Certificate{
			KeyId: "llama",
			Permissions: ssh.Permissions{
				Extensions: map[string]string{
					teleport.CertExtensionDeviceID:           ext.DeviceID,
					teleport.CertExtensionDeviceAssetTag:     ext.AssetTag,
					teleport.CertExtensionDeviceCredentialID: ext.CredentialID,
				},
			},
		})
	})
}

func runVerifyUserTest(t *testing.T, method string, verify func(dt *types.DeviceTrust, ext *tlsca.DeviceExtensions) error) {
	assertNoErr := func(t *testing.T, err error) {
		assert.NoError(t, err, "%v mismatch", method)
	}
	assertDeniedErr := func(t *testing.T, err error) {
		assert.ErrorContains(t, err, "unauthorized device", "%v mismatch", method)
		assert.True(t, trace.IsAccessDenied(err), "%v returned an error other than trace.AccessDeniedError: %T", method, err)
	}

	userWithoutExtensions := &tlsca.DeviceExtensions{}
	userWithExtensions := &tlsca.DeviceExtensions{
		DeviceID:     "deviceid1",
		AssetTag:     "assettag1",
		CredentialID: "credentialid1",
	}

	tests := []struct {
		name      string
		buildType string
		dt        *types.DeviceTrust
		ext       *tlsca.DeviceExtensions
		assertErr func(t *testing.T, err error)
	}{
		{
			name:      "OSS dt=nil",
			buildType: modules.BuildOSS,
			dt:        nil, // OK, config absent.
			ext:       userWithoutExtensions,
			assertErr: assertNoErr,
		},
		{
			name:      "OSS mode=off",
			buildType: modules.BuildOSS,
			dt: &types.DeviceTrust{
				Mode: constants.DeviceTrustModeOff, // Valid for OSS.
			},
			ext:       userWithoutExtensions,
			assertErr: assertNoErr,
		},
		{
			name:      "OSS mode never enforced",
			buildType: modules.BuildOSS,
			dt: &types.DeviceTrust{
				Mode: constants.DeviceTrustModeRequired, // Invalid for OSS, treated as "off".
			},
			ext:       userWithoutExtensions,
			assertErr: assertNoErr,
		},
		{
			name:      "Enterprise mode=off",
			buildType: modules.BuildEnterprise,
			dt: &types.DeviceTrust{
				Mode: constants.DeviceTrustModeOff,
			},
			ext:       userWithoutExtensions,
			assertErr: assertNoErr,
		},
		{
			name:      "Enterprise mode=optional without extensions",
			buildType: modules.BuildEnterprise,
			dt: &types.DeviceTrust{
				Mode: constants.DeviceTrustModeOptional,
			},
			ext:       userWithoutExtensions,
			assertErr: assertNoErr,
		},
		{
			name:      "Enterprise mode=optional with extensions",
			buildType: modules.BuildEnterprise,
			dt: &types.DeviceTrust{
				Mode: constants.DeviceTrustModeOptional,
			},
			ext:       userWithExtensions, // Happens if the device is enrolled.
			assertErr: assertNoErr,
		},
		{
			name:      "nok: Enterprise mode=required without extensions",
			buildType: modules.BuildEnterprise,
			dt: &types.DeviceTrust{
				Mode: constants.DeviceTrustModeRequired,
			},
			ext:       userWithoutExtensions,
			assertErr: assertDeniedErr,
		},
		{
			name:      "Enterprise mode=required with extensions",
			buildType: modules.BuildEnterprise,
			dt: &types.DeviceTrust{
				Mode: constants.DeviceTrustModeRequired,
			},
			ext:       userWithExtensions,
			assertErr: assertNoErr,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules.SetTestModules(t, &modules.TestModules{
				TestBuildType: test.buildType,
			})

			test.assertErr(t, verify(test.dt, test.ext))
		})
	}
}
