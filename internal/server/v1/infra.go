package server

import (
	"io/ioutil"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
)

var (
	logger grpclog.LoggerV2
	db     *sqlx.DB
)

func bootstrapInfra() {
	logger = grpclog.NewLoggerV2WithVerbosity(os.Stdout, ioutil.Discard, ioutil.Discard, viper.GetInt("log.level"))

	var err error

	db, err = sqlx.Connect("postgres", viper.GetString("database.url"))
	if err != nil {
		logger.Fatalf("[db] cannot connect to the database: %s", err.Error())
	}
}
