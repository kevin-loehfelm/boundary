// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package warning

import (
	"context"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hashicorp/boundary/internal/observability/event"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/hashicorp/boundary/internal/errors"
	pbwarnings "github.com/hashicorp/boundary/internal/gen/controller/api"
)

type warningKey int

var (
	warnerContextkey warningKey
	warningHeader    = "x-boundary-warning"
)

type Warner struct {
	fieldWarnings map[string][]string
}

func (w *Warner) AddFieldWarning(field, warning string) {
	w.fieldWarnings[field] = append(w.fieldWarnings[field], warning)
}

func FromContext(ctx context.Context) (*Warner, bool) {
	w, ok := ctx.Value(warnerContextkey).(*Warner)
	return w, ok
}

func newContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, warnerContextkey, &Warner{fieldWarnings: make(map[string][]string)})
}

func convertToGrpcHeaders(ctx context.Context) error {
	const op = "warning.convertToGrpcHeaders"
	w, ok := FromContext(ctx)
	if !ok {
		return errors.New(ctx, errors.InvalidParameter, op, "context doesn't have warner")
	}
	if len(w.fieldWarnings) == 0 {
		return nil
	}

	pbWar := &pbwarnings.Warning{}
	for k, v := range w.fieldWarnings {
		pbWar.RequestFields = append(pbWar.RequestFields, &pbwarnings.FieldWarning{
			Name:        k,
			Description: strings.Join(v, ", "),
		})
	}
	var buf []byte
	var err error
	if buf, err = protojson.Marshal(pbWar); err != nil {
		return errors.Wrap(ctx, err, op, errors.WithMsg("unable to marshal warnings"))
	}
	if err := grpc.SetHeader(ctx, metadata.Pairs(warningHeader, string(buf))); err != nil {
		return errors.Wrap(ctx, err, op, errors.WithMsg("unable to set warning grpc header"))
	}
	return nil
}

func OutgoingHeaderMatcher() runtime.HeaderMatcherFunc {
	return func(s string) (string, bool) {
		if s == warningHeader {
			return warningHeader, true
		}
		return "", false
	}
}

// GrpcInterceptor intercepts warnings as reported by the handlers and populates
// them in a specific header.
func GrpcInterceptor(outerCtx context.Context) grpc.UnaryServerInterceptor {
	const op = "controller.warningInterceptor"
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx = newContext(ctx)
		h, handlerErr := handler(ctx, req)
		if err := convertToGrpcHeaders(ctx); err != nil {
			event.WriteError(outerCtx, op, err)
		}
		return h, handlerErr
	}
}
