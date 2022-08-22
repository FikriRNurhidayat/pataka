package server

import (
	"context"
	"fmt"
	"net"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	GRPC_PORT        string
	GRPC_ENDPOINT    string
	GATEWAY_PORT     string
	GATEWAY_ENDPOINT string
)

func Serve() {
	bootstrapInfra()
	bootstrapServers()

	GRPC_PORT = viper.GetString("grpc.port")
	GRPC_ENDPOINT = fmt.Sprintf(":%s", GRPC_PORT)
	GATEWAY_PORT = viper.GetString("gateway.port")
	GATEWAY_ENDPOINT = fmt.Sprintf(":%s", GATEWAY_PORT)

	// Set logger
	grpclog.SetLoggerV2(logger)

	// Create context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Initiate TCP Listener
	lis, err := net.Listen("tcp", GRPC_ENDPOINT)
	if err != nil {
		logger.Fatalf("[net] failed to initialize TCP connection: %s", err.Error())
	}

	s := grpc.NewServer()
	s = registerServer(s)

	// Register and run GRPC Server
	go s.Serve(lis)

	// Run gateway
	if err := createGateway(ctx).ListenAndServe(); err != nil {
		logger.Fatalf("[gateway] failed to start gateway server: %s", err.Error())
	}
}
