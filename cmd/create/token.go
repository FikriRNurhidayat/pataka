/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"
	"os"

	"github.com/fikrirnurhidayat/ffgo/internal/app/authentication"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// tokenCmd represents the createToken command
var TokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate new access token",
	Long:  `In order to access Pataka API, you need an access token.`,
	Run: func(cmd *cobra.Command, args []string) {
		as := authentication.NewAuthenticaticationService(viper.GetString("admin.secret"))
		token, err := as.CreateToken()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(token)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.BindEnv("admin.secret", "ADMIN_SECRET")
}

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

	viper.SetEnvPrefix("pataka")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
