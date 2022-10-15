/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fikrirnurhidayat/ffgo/internal/app/authentication"
	"github.com/fikrirnurhidayat/ffgo/internal/command/create"
	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		authenticationService := authentication.New("Rahasia")
		createTokenCmd := create.NewCreateTokenCmd(authenticationService)

		scopes, _ := cmd.Flags().GetStringSlice("scopes")
		token, _ := createTokenCmd.Call(scopes...)

		fmt.Println(token)
	},
}

func init() {
	createCmd.AddCommand(tokenCmd)
	tokenCmd.Flags().StringSlice("scopes", []string{}, "Define token scopes.")
}
