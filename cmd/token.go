/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fikrirnurhidayat/ffgo/internal/auth"
	"github.com/fikrirnurhidayat/ffgo/internal/command/create"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		authenticationService := authentication.New(viper.GetString("secretKey"))
		createTokenCmd := create.NewCreateTokenCmd(authenticationService)

		scopes, _ := cmd.Flags().GetStringSlice("scopes")
		token, _ := createTokenCmd.Call(scopes...)

		fmt.Println(token)
	},
}

func init() {
	createCmd.AddCommand(tokenCmd)
	tokenCmd.Flags().StringSlice("scopes", []string{}, "Define token scopes.")
	viper.BindEnv("secretKey", "SECRET_KEY")
}
