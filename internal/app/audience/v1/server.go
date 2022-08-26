package audience

import (
	"github.com/fikrirnurhidayat/ffgo/internal/app/feature/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/driver"
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
	db                         driver.DB
}

type ServerOpts func(*Server)

func NewServer(opts ...ServerOpts) audiencev1.AudienceServiceServer {
	s := new(Server)

	for _, set := range opts {
		set(s)
	}

	authentication := auth.New(viper.GetString("admin.secret"))
	audienceRepository := NewPostgresRepository(s.db, s.logger)
	featureRepository := feature.NewPostgresRepository(s.db, s.logger)
	s.createAudienceService = NewCreateAudienceService(authentication, audienceRepository, featureRepository, s.logger)
	s.listAudiencesService = NewListAudiencesService(audienceRepository, s.logger, 1, 10)
	s.getAudienceService = NewGetAudienceService(audienceRepository, s.logger)
	s.updateAudienceService = NewUpdateAudienceService(authentication, audienceRepository, s.logger)
	s.deleteAudienceService = NewDeleteAudienceService(authentication, audienceRepository, s.logger)

	return s
}

func WithLogger(logger grpclog.LoggerV2) ServerOpts {
	return func(fs *Server) {
		fs.logger = logger
	}
}

func WithDB(db driver.DB) ServerOpts {
	return func(fs *Server) {
		fs.db = db
	}
}
