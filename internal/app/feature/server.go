package feature

import (
	"github.com/fikrirnurhidayat/ffgo/gen/proto/go/featureflag/v1"
	"google.golang.org/grpc/grpclog"
)

type FeatureServer struct {
	featureflag.UnimplementedFeatureServiceServer
	Create Createable
	Delete Deletable
	Update Updatable
	Get    Getable
	List   Listable
	Logger grpclog.LoggerV2
}

type FeatureServerOpts func(*FeatureServer)

func NewFeatureServer(opts ...FeatureServerOpts) featureflag.FeatureServiceServer {
	server := new(FeatureServer)

	for _, set := range opts {
		set(server)
	}

	return server
}

func WithCreator(createFunc Createable) FeatureServerOpts {
	return func(fs *FeatureServer) {
		fs.Create = createFunc
	}
}

func WithDeletor(deleteFunc Deletable) FeatureServerOpts {
	return func(fs *FeatureServer) {
		fs.Delete = deleteFunc
	}
}

func WithUpdater(updateFunc Updatable) FeatureServerOpts {
	return func(fs *FeatureServer) {
		fs.Update = updateFunc
	}
}

func WithGetter(getFunc Getable) FeatureServerOpts {
	return func(fs *FeatureServer) {
		fs.Get = getFunc
	}
}

func WithLister(listFunc Listable) FeatureServerOpts {
	return func(fs *FeatureServer) {
		fs.List = listFunc
	}
}

func WithLogger(logger grpclog.LoggerV2) FeatureServerOpts {
	return func(fs *FeatureServer) {
		fs.Logger = logger
	}
}
