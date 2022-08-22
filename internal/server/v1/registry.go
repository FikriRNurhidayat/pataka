package server

import (
	"context"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func registerServer(s *grpc.Server) *grpc.Server {
	featurev1.RegisterFeatureServiceServer(s, featureServer)
	audiencev1.RegisterAudienceServiceServer(s, audienceServer)

	return s
}

func registerEndpoint(ctx context.Context, mux *runtime.ServeMux, dialOptions []grpc.DialOption) *runtime.ServeMux {
	featurev1.RegisterFeatureServiceHandlerFromEndpoint(ctx, mux, GRPC_ENDPOINT, dialOptions)
	audiencev1.RegisterAudienceServiceHandlerFromEndpoint(ctx, mux, GRPC_ENDPOINT, dialOptions)

	return mux
}
