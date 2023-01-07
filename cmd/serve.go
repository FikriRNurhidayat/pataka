/*
Copyright © 2022 FIKRI RAHMAT NURHIDAYAT <fikrirnurhidayat@gmail.com>
*/
package cmd

import (
	"github.com/fikrirnurhidayat/ffgo/internal/server/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run feature flags server",
	Long:  `Serve feature flags server, make sure you have database properly setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Serve()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().String("grpc-port", "8080", "grpc port to listen to.")
	serveCmd.Flags().String("gateway-port", "8081", "gateway port to listen to.")

	viper.BindPFlag("bind", serveCmd.Flags().Lookup("bind"))
	viper.BindPFlag("grpc.port", serveCmd.Flags().Lookup("grpc-port"))
	viper.BindPFlag("gateway.port", serveCmd.Flags().Lookup("gateway-port"))
	viper.BindEnv("database.url", "PATAKA_DATABASE_URL")
	viper.BindEnv("database.pool", "PATAKA_DATABASE_POOL")
	viper.BindEnv("redis.url", "PATAKA_REDIS_URL")
	viper.BindEnv("log.level", "PATAKA_LOG_LEVEL")
	viper.BindEnv("bind", "PATAKA_BIND")
	viper.BindEnv("secretKey", "PATAKA_SECRET_KEY")
}
