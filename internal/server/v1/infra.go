package server

import (
	"io/ioutil"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
)

var (
	logger      grpclog.LoggerV2
	db          *sqlx.DB
	redisClient redis.UniversalClient
)

func bootstrapInfra() {
	logger = grpclog.NewLoggerV2WithVerbosity(os.Stdout, ioutil.Discard, ioutil.Discard, viper.GetInt("log.level"))

	var err error

	db, err = sqlx.Connect("postgres", viper.GetString("database.url"))
	if err != nil {
		logger.Fatalf("[db] cannot connect to the database: %s", err.Error())
	}

	db.SetMaxIdleConns(viper.GetInt("database.pool"))
	db.SetMaxOpenConns(viper.GetInt("database.pool"))

	redisUrl := viper.GetString("redis.url")
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		logger.Infoln(opt)
		logger.Fatalf("[cache] cannot connect to the cache: %s", err.Error())
	}

	redisClient = redis.NewClient(opt)
}
