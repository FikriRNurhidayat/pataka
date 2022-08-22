package audience

import (
	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"github.com/fikrirnurhidayat/ffgo/internal/driver/databasesql"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"

	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

type Server struct {
	audiencev1.UnimplementedAudienceServiceServer
	BulkCreate BulkCreateable
	Create     Createable
	Delete     Deletable
	Update     Updatable
	Get        Getable
	List       Listable
	Logger     grpclog.LoggerV2
	DB         databasesql.DB
}

type ServerOpts func(*Server)

func NewServer(opts ...ServerOpts) audiencev1.AudienceServiceServer {
	s := new(Server)

	for _, set := range opts {
		set(s)
	}

	authentication := auth.New(viper.GetString("admin.secret"))
	audienceRepository := NewPostgresRepository(s.DB, s.Logger)
	s.Create = NewCreateAudienceService(authentication, audienceRepository, s.Logger)
	s.List = NewListAudiencesService(audienceRepository, s.Logger, 1, 10)
	s.Get = NewGetAudienceService(audienceRepository, s.Logger)
	s.Update = NewUpdateAudienceService(authentication, audienceRepository, s.Logger)
	s.Delete = NewDeleteAudienceService(authentication, audienceRepository, s.Logger)

	return s
}

func WithLogger(logger grpclog.LoggerV2) ServerOpts {
	return func(fs *Server) {
		fs.Logger = logger
	}
}

func WithDB(db databasesql.DB) ServerOpts {
	return func(fs *Server) {
		fs.DB = db
	}
}
