package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func createGateway(ctx context.Context) *http.Server {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				AllowPartial:    true,
				UseProtoNames:   true,
				UseEnumNumbers:  false,
				EmitUnpopulated: false,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				AllowPartial:   true,
				DiscardUnknown: true,
			},
		}),
	)

	registerEndpoint(ctx, mux, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})

	return &http.Server{
		Addr:    GATEWAY_ENDPOINT,
		Handler: mux,
	}
}
