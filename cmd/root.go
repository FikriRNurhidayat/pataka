/*
Copyright © 2022 FIKRI RAHMAT NURHIDAYAT <fikrirnurhidayat@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pataka",
	Short: "Minimal feature flag server.",
	Long:  `Feature flag server that are pluggable with your existing applications`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.ffgo.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ffgo" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ffgo")
	}

	viper.SetEnvPrefix("PATAKA")
	viper.BindEnv("GRPC_PORT")
	viper.BindEnv("GATEWAY_PORT")
	viper.BindEnv("DATABASE_URL")
	viper.BindEnv("REDIS_URL")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
