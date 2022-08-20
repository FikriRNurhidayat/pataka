package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/app/authentication"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	GRPC_PORT             string
	GRPC_ENDPOINT         string
	GATEWAY_PORT          string
	GATEWAY_ENDPOINT      string
	logger                grpclog.LoggerV2
	db                    *sqlx.DB
	authenticationService authentication.Authenticatable
)

func initAuthenticationService() {
	authenticationService = authentication.NewAuthenticaticationService(viper.GetString("admin.secret"))
}

func initInfra() {
	logger = grpclog.NewLoggerV2WithVerbosity(os.Stdout, ioutil.Discard, ioutil.Discard, viper.GetInt("log.level"))

	var err error

	db, err = sqlx.Connect("postgres", viper.GetString("database.url"))
	if err != nil {
		logger.Fatalf("[db] cannot connect to the database: %s", err.Error())
	}
}

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

	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register feature service handler from endpoint
	featureflag.RegisterFeatureServiceHandlerFromEndpoint(ctx, mux, GRPC_ENDPOINT, dialOptions)

	return &http.Server{
		Addr:    GATEWAY_ENDPOINT,
		Handler: mux,
	}
}

func Serve() {
	initInfra()
	initAuthenticationService()
	initFeatureServer()

	GRPC_PORT = viper.GetString("grpc.port")
	GRPC_ENDPOINT = fmt.Sprintf(":%s", GRPC_PORT)
	GATEWAY_PORT = viper.GetString("gateway.port")
	GATEWAY_ENDPOINT = fmt.Sprintf(":%s", GATEWAY_PORT)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Initiate TCP Listener
	lis, err := net.Listen("tcp", GRPC_ENDPOINT)
	if err != nil {
		logger.Fatalf("[net] failed to initialize TCP connection: %s", err.Error())
	}

	server := grpc.NewServer()
	grpclog.SetLoggerV2(logger)

	// Register featureServer
	featureflag.RegisterFeatureServiceServer(server, featureServer)

	go server.Serve(lis)

	// Gateway
	gateway := createGateway(ctx)
	if err := gateway.ListenAndServe(); err != nil {
		logger.Fatalf("[gateway] failed to start gateway server: %s", err.Error())
	}
}
