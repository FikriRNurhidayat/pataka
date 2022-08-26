package feature

import (
	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
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
	featureRepository    domain.FeatureRepository
	logger               grpclog.LoggerV2
}

type ServerOpts func(*Server)

func NewServer(opts ...ServerOpts) featurev1.FeatureServiceServer {
	s := new(Server)

	for _, set := range opts {
		set(s)
	}

	authentication := auth.New(viper.GetString("admin.secret"))
	s.createFeatureService = NewCreateFeatureService(authentication, s.featureRepository, s.logger)
	s.listFeaturesService = NewListFeaturesService(s.featureRepository, s.logger, 1, 10)
	s.getFeatureService = NewGetFeatureService(s.featureRepository, s.logger)
	s.updateFeatureService = NewUpdateFeatureService(authentication, s.featureRepository, s.logger)
	s.deleteFeatureService = NewDeleteFeatureService(authentication, s.featureRepository, s.logger)

	return s
}

func WithLogger(logger grpclog.LoggerV2) ServerOpts {
	return func(fs *Server) {
		fs.logger = logger
	}
}

func WithFeatureRepository(featureRepository domain.FeatureRepository) ServerOpts {
	return func(fs *Server) {
		fs.featureRepository = featureRepository
	}
}
