// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package warning

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hashicorp/boundary/globals"
	pbwarnings "github.com/hashicorp/boundary/internal/gen/controller/api"
	opsservices "github.com/hashicorp/boundary/internal/gen/ops/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestFromContext(t *testing.T) {
	t.Run("no warning on context", func(t *testing.T) {
		w, ok := FromContext(context.Background())
		require.False(t, ok)
		require.Nil(t, w)
	})

	t.Run("empty warning on context", func(t *testing.T) {
		ctx := newContext(context.Background())
		w, ok := FromContext(ctx)
		require.True(t, ok)
		require.NotNil(t, w)
	})

	t.Run("warning added after added to context", func(t *testing.T) {
		ctx := newContext(context.Background())
		w, ok := FromContext(ctx)
		require.True(t, ok)
		require.NotNil(t, w)
		w.AddFieldWarning("test_field", "this is a test")

		newW, newOk := FromContext(ctx)
		require.True(t, newOk)
		require.NotNil(t, newW)
		assert.Equal(t, &Warner{fieldWarnings: []*pbwarnings.FieldWarning{
			{
				Name:        "test_field",
				Description: "this is a test",
			},
		}}, newW)
	})
}

func TestGrpcGatwayWiring(t *testing.T) {
	ctx := context.Background()
	fieldWarnings := []*pbwarnings.FieldWarning{
		{
			Name:        "test_field_1",
			Description: "test warning description 1",
		},
		{
			Name:        "test_field_2",
			Description: "test warning description 2",
		},
	}

	want, err := protojson.Marshal(&pbwarnings.Warning{
		RequestFields: fieldWarnings,
	})

	grpcSrv := grpc.NewServer(grpc.UnaryInterceptor(GrpcInterceptor(ctx)))
	opsservices.RegisterHealthServiceServer(grpcSrv, &fakeService{
		addWarnFunc: func(ctx context.Context) {
			g, ok := FromContext(ctx)
			require.True(t, ok)
			for _, w := range fieldWarnings {
				g.AddFieldWarning(w.GetName(), w.GetDescription())
			}
		},
	})

	l := bufconn.Listen(int(globals.DefaultMaxRequestSize))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcSrv.Serve(l)
	}()

	gwMux := runtime.NewServeMux(
		runtime.WithOutgoingHeaderMatcher(OutgoingHeaderMatcher()),
	)
	require.NoError(t, opsservices.RegisterHealthServiceHandlerFromEndpoint(ctx, gwMux, "", []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return l.Dial()
		}),
	}))

	mux := http.NewServeMux()
	mux.Handle("/health", gwMux)

	httpSrv := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		IdleTimeout:       5 * time.Minute,
	}

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpSrv.Serve(lis)
	}()

	resp, err := http.Get(fmt.Sprintf("http://%s/health", lis.Addr().String()))
	require.NoError(t, err)
	got := resp.Header.Get(warningHeader)

	require.NoError(t, err)
	assert.Equal(t, string(want), got)

	assert.NoError(t, httpSrv.Shutdown(ctx))
	grpcSrv.GracefulStop()
	wg.Wait()
}

// fakeService is made to
type fakeService struct {
	opsservices.UnimplementedHealthServiceServer
	addWarnFunc func(context.Context)
}

func (f fakeService) GetHealth(ctx context.Context, request *opsservices.GetHealthRequest) (*opsservices.GetHealthResponse, error) {
	f.addWarnFunc(ctx)
	return &opsservices.GetHealthResponse{}, nil
}
