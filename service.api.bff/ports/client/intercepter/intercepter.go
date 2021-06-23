package intercepter

import (
	"context"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func WithUnary(f grpc.UnaryClientInterceptor) grpc.DialOption {
	return grpc.WithUnaryInterceptor(f)
}

// WithAuth serves as an grpc client intercepter appending the grpc context with the
// JWT user token for authentication on service level
func WithAuth(ctx context.Context, method string, req interface{}, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
) error {
	accessToken := ctx_value.GetString(ctx, "authorization")
	meta := metadata.Pairs("authorization", accessToken)
	newCtx := metadata.NewOutgoingContext(ctx, meta)
	return invoker(newCtx, method, req, reply, cc, opts...)
}
