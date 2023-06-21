// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package warning

import (
	"context"

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
	warningHeader    = "x-boundary-warnings"
)

type warner struct {
	fieldWarnings []*pbwarnings.FieldWarning
}

func (w *warner) addFieldWarning(field, warning string) {
	w.fieldWarnings = append(w.fieldWarnings, &pbwarnings.FieldWarning{
		Name:        field,
		Description: warning,
	})
}

func ForField(ctx context.Context, field, warning string) error {
	const op = "warning.ForField"
	w, ok := ctx.Value(warnerContextkey).(*warner)
	if !ok {
		return errors.New(ctx, errors.InvalidParameter, op, "context doesn't contain warning functionality")
	}
	w.addFieldWarning(field, warning)
	return nil
}

func newContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, warnerContextkey, &warner{})
}

func convertToGrpcHeaders(ctx context.Context) error {
	const op = "warning.convertToGrpcHeaders"
	w, ok := ctx.Value(warnerContextkey).(*warner)
	if !ok {
		return errors.New(ctx, errors.InvalidParameter, op, "context doesn't have warner")
	}
	if len(w.fieldWarnings) == 0 {
		return nil
	}

	pbWar := &pbwarnings.Warning{RequestFields: w.fieldWarnings}
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

// OutgoingHeaderMatcher provides a runtime.HeaderMatcherFunc that can be used
// as an option when creating a new grpc gateway server muxer, and specifies
// the boundary warning headers which can be forwarded on to the requesting client.
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
