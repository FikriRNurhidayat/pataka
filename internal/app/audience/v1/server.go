package audience

import (
	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

type Server struct {
	audiencev1.UnimplementedAudienceServiceServer
	bulkCreateAudiencesService domain.AudienceBulkCreatable
	createAudienceService      domain.AudienceCreatable
	deleteAudienceService      domain.AudienceDeletable
	updateAudienceService      domain.AudienceUpdateable
	getAudienceService         domain.AudienceGetable
	listAudiencesService       domain.AudienceListable
	logger                     grpclog.LoggerV2
	audienceRepository         domain.AudienceRepository
	featureRepository          domain.FeatureRepository
}

type ServerOpts func(*Server)

func NewServer(opts ...ServerOpts) audiencev1.AudienceServiceServer {
	s := new(Server)

	for _, set := range opts {
		set(s)
	}

	authentication := auth.New(viper.GetString("admin.secret"))
	s.createAudienceService = NewCreateAudienceService(authentication, s.audienceRepository, s.featureRepository, s.logger)
	s.listAudiencesService = NewListAudiencesService(s.audienceRepository, s.logger, 1, 10)
	s.getAudienceService = NewGetAudienceService(s.audienceRepository, s.logger)
	s.updateAudienceService = NewUpdateAudienceService(authentication, s.audienceRepository, s.logger)
	s.deleteAudienceService = NewDeleteAudienceService(authentication, s.audienceRepository, s.logger)

	return s
}

func WithLogger(logger grpclog.LoggerV2) ServerOpts {
	return func(fs *Server) {
		fs.logger = logger
	}
}

func WithAudienceRepository(audienceRepository domain.AudienceRepository) ServerOpts {
	return func(fs *Server) {
		fs.audienceRepository = audienceRepository
	}
}

func WithFeatureRepository(featureRepository domain.FeatureRepository) ServerOpts {
	return func(fs *Server) {
		fs.featureRepository = featureRepository
	}
}
