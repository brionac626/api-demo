package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var configPath string

var apiCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a http server for article api service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("app server test")
	},
}

var boAPICmd = &cobra.Command{
	Use:   "back-office",
	Short: "Start the http server for the back office service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("app back-office test")
	},
}
