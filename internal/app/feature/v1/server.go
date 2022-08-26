package feature

import (
	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/driver"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"

	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

type Server struct {
	featurev1.UnimplementedFeatureServiceServer
	createFeatureService domain.FeatureCreateable
	deleteFeatureService domain.FeatureDeletable
	updateFeatureService domain.FeatureUpdatable
	getFeatureService    domain.FeatureGetable
	listFeaturesService  domain.FeaturesListable
	Logger               grpclog.LoggerV2
	DB                   driver.DB
}

type ServerOpts func(*Server)

func NewServer(opts ...ServerOpts) featurev1.FeatureServiceServer {
	s := new(Server)

	for _, set := range opts {
		set(s)
	}

	authentication := auth.New(viper.GetString("admin.secret"))
	featureRepository := NewPostgresRepository(s.DB, s.Logger)
	s.createFeatureService = NewCreateFeatureService(authentication, featureRepository, s.Logger)
	s.listFeaturesService = NewListFeaturesService(featureRepository, s.Logger, 1, 10)
	s.getFeatureService = NewGetFeatureService(featureRepository, s.Logger)
	s.updateFeatureService = NewUpdateFeatureService(authentication, featureRepository, s.Logger)
	s.deleteFeatureService = NewDeleteFeatureService(authentication, featureRepository, s.Logger)

	return s
}

func WithLogger(logger grpclog.LoggerV2) ServerOpts {
	return func(fs *Server) {
		fs.Logger = logger
	}
}

func WithDB(db driver.DB) ServerOpts {
	return func(fs *Server) {
		fs.DB = db
	}
}
