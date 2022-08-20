package feature

import (
	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
)

type Server struct {
	featureflag.UnimplementedFeatureServiceServer
	Create Createable
	Delete Deletable
	Update Updatable
	Get    Getable
	List   Listable
	Logger grpclog.LoggerV2
	DB     DB
}

type ServerOpts func(*Server)

func NewServer(opts ...ServerOpts) featureflag.FeatureServiceServer {
	s := new(Server)

	for _, set := range opts {
		set(s)
	}

	authentication := auth.New(viper.GetString("admin.secret"))
	featureRepository := NewPostgresFeatureRepository(s.DB, s.Logger)
	s.Create = NewCreateFeatureService(authentication, featureRepository, s.Logger)
	s.List = NewListFeaturesService(featureRepository, s.Logger, 1, 10)
	s.Get = NewGetFeatureService(featureRepository, s.Logger)
	s.Update = NewUpdateFeatureService(authentication, featureRepository, s.Logger)
	s.Delete = NewDeleteFeatureService(authentication, featureRepository, s.Logger)

	return s
}

func WithLogger(logger grpclog.LoggerV2) ServerOpts {
	return func(fs *Server) {
		fs.Logger = logger
	}
}

func WithDB(db DB) ServerOpts {
	return func(fs *Server) {
		fs.DB = db
	}
}
