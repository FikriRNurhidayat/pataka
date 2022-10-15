package audience

import (
	"github.com/fikrirnurhidayat/ffgo/internal/app/authentication"
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

type Server struct {
	audiencev1.UnimplementedAudienceServiceServer
	authenticationService      domain.Authenticatable
	bulkCreateAudiencesService domain.AudienceBulkCreatable
	createAudienceService      domain.AudienceCreatable
	deleteAudienceService      domain.AudienceDeletable
	updateAudienceService      domain.AudienceUpdateable
	getAudienceService         domain.AudienceGetable
	listAudiencesService       domain.AudienceListable
	logger                     grpclog.LoggerV2
	unitOfWork                 domain.UnitOfWork
	audienceRepository         domain.AudienceRepository
}

type ServerSetter func(*Server)

func NewServer(opts ...ServerSetter) audiencev1.AudienceServiceServer {
	s := new(Server)

	for _, set := range opts {
		set(s)
	}

	secretKey := viper.GetString("secretKey")
	s.authenticationService = authentication.New(secretKey)
	s.createAudienceService = NewCreateAudienceService(s.unitOfWork, s.logger)
	s.listAudiencesService = NewListAudiencesService(s.audienceRepository, s.logger)
	s.getAudienceService = NewGetAudienceService(s.audienceRepository, s.logger)
	s.updateAudienceService = NewUpdateAudienceService(s.unitOfWork, s.logger)
	s.deleteAudienceService = NewDeleteAudienceService(s.unitOfWork, s.logger)

	return s
}

func WithLogger(logger grpclog.LoggerV2) ServerSetter {
	return func(fs *Server) {
		fs.logger = logger
	}
}

func WithUnitOfWork(unitOfWork domain.UnitOfWork) ServerSetter {
	return func(fs *Server) {
		fs.unitOfWork = unitOfWork
	}
}

func WithAudienceRepository(audienceRepository domain.AudienceRepository) ServerSetter {
	return func(fs *Server) {
		fs.audienceRepository = audienceRepository
	}
}
