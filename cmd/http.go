package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var configPath string

var httpCmd = &cobra.Command{
	Use: "http",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("app test")
	},
}
