package feature

import (
	"github.com/fikrirnurhidayat/ffgo/internal/app/authentication"
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"

	featurev1 "github.com/fikrirnurhidayat/ffgo/protobuf/feature/v1"
)

type Server struct {
	featurev1.UnimplementedFeatureServiceServer
	authenticationService domain.Authenticatable
	createFeatureService  domain.FeatureCreateable
	deleteFeatureService  domain.FeatureDeletable
	updateFeatureService  domain.FeatureUpdatable
	getFeatureService     domain.FeatureGetable
	listFeaturesService   domain.FeaturesListable
	featureRepository     domain.FeatureRepository
	unitOfWork            domain.UnitOfWork
	logger                grpclog.LoggerV2
}

type ServerOpts func(*Server)

func NewServer(opts ...ServerOpts) featurev1.FeatureServiceServer {
	s := new(Server)

	for _, set := range opts {
		set(s)
	}

	secretKey := viper.GetString("secretKey")
	s.authenticationService = authentication.New(secretKey)
	s.createFeatureService = NewCreateFeatureService(s.unitOfWork, s.logger)
	s.listFeaturesService = NewListFeaturesService(s.featureRepository, s.logger)
	s.getFeatureService = NewGetFeatureService(s.featureRepository, s.logger)
	s.updateFeatureService = NewUpdateFeatureService(s.unitOfWork, s.logger)
	s.deleteFeatureService = NewDeleteFeatureService(s.unitOfWork, s.logger)

	return s
}

func WithLogger(logger grpclog.LoggerV2) ServerOpts {
	return func(fs *Server) {
		fs.logger = logger
	}
}

func WithUnitOfWork(unitOfWork domain.UnitOfWork) ServerOpts {
	return func(fs *Server) {
		fs.unitOfWork = unitOfWork
	}
}

func WithFeatureRepository(featureRepository domain.FeatureRepository) ServerOpts {
	return func(fs *Server) {
		fs.featureRepository = featureRepository
	}
}
