/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
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

package machineidv1

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"github.com/google/go-cmp/cmp"
	headerv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	machineidv1pb "github.com/gravitational/teleport/api/gen/proto/go/teleport/machineid/v1"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/backend/memory"
	"github.com/gravitational/teleport/lib/fixtures"
	"github.com/gravitational/teleport/lib/services/local"
	"github.com/gravitational/teleport/lib/utils"
	"github.com/jonboulle/clockwork"
	"github.com/spiffe/go-spiffe/v2/bundle/spiffebundle"
	"github.com/spiffe/go-spiffe/v2/federation"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSPIFFEFederationSyncer(t *testing.T) {
	t.Fatalf("Test not implemented")
}

func TestSPIFFEFederationSyncer_syncFederation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	logger := utils.NewSlogLoggerForTests()
	clock := clockwork.NewFakeClock()
	backend, err := memory.New(memory.Config{})
	require.NoError(t, err)

	store, err := local.NewSPIFFEFederationService(backend)
	require.NoError(t, err)
	eventsSvc := local.NewEventsService(backend)

	td := spiffeid.RequireTrustDomainFromString("example.com")
	bundle := spiffebundle.New(td)
	b, _ := pem.Decode([]byte(fixtures.TLSCACertPEM))
	cert, err := x509.ParseCertificate(b.Bytes)
	if !assert.NoError(t, err) {
		return
	}
	bundle.AddX509Authority(cert)
	bundle.SetRefreshHint(time.Minute * 12)
	marshaledBundle, err := bundle.Marshal()
	require.NoError(t, err)

	// Implement a fake SPIFFE Federation bundle endpoint
	testSrv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h, err := federation.NewHandler(td, bundle)
		if !assert.NoError(t, err) {
			return
		}
		h.ServeHTTP(w, r)
	}))
	caPool := x509.NewCertPool()
	caPool.AddCert(testSrv.Certificate())

	syncer, err := NewSPIFFEFederationSyncer(SPIFFEFederationSyncerConfig{
		Backend:       backend,
		Store:         store,
		EventsWatcher: eventsSvc,
		Clock:         clock,
		Logger:        logger,
		SPIFFEFetchOptions: []federation.FetchOption{
			federation.WithWebPKIRoots(caPool),
		},
	})
	require.NoError(t, err)

	in := &machineidv1pb.SPIFFEFederation{
		Kind:    types.KindSPIFFEFederation,
		Version: types.V1,
		Metadata: &headerv1.Metadata{
			Name: "example.com",
		},
		Spec: &machineidv1pb.SPIFFEFederationSpec{
			BundleSource: &machineidv1pb.SPIFFEFederationBundleSource{
				HttpsWeb: &machineidv1pb.SPIFFEFederationBundleSourceHTTPSWeb{
					BundleEndpointUrl: testSrv.URL,
				},
			},
		},
	}
	created, err := store.CreateSPIFFEFederation(ctx, in)
	require.NoError(t, err)

	// TODO: Add test for static bundle, change and no change
	// TODO: Add test for already synced recently (e.g sync not necessary)

	out, err := syncer.syncFederation(ctx, "example.com")
	require.NoError(t, err)

	got, err := store.GetSPIFFEFederation(ctx, "example.com")
	require.NoError(t, err)
	// Require that the persisted resource equals the resource output by syncFederation
	require.Empty(t, cmp.Diff(out, got, protocmp.Transform()))
	// Check that some update as occurred (as indicated by the revision)
	require.NotEqual(t, created.Metadata.Revision, got.Metadata.Revision)
	// Check that the expected status fields have been set...
	require.NotNil(t, got.Status)
	wantStatus := &machineidv1pb.SPIFFEFederationStatus{
		CurrentBundleSyncedAt:   timestamppb.New(clock.Now().UTC()),
		CurrentBundleSyncedFrom: proto.Clone(created).(*machineidv1pb.SPIFFEFederation).Spec.BundleSource,
		NextSyncAt:              timestamppb.New(clock.Now().UTC().Add(time.Minute * 12)),
		CurrentBundle:           string(marshaledBundle),
	}
	require.Empty(t, cmp.Diff(got.Status, wantStatus, protocmp.Transform()))
}
