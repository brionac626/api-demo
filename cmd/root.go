/*
Copyright © 2024 HSINE KUEI CHIU <theone1632@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	SilenceUsage: true,
	Short:        "A api server demo",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

func init() {
	apiCmd.Flags().StringVarP(&configPath, "config", "c", "./deployment/local/config.yaml", "configuration file path")
	boAPICmd.Flags().StringVarP(&configPath, "config", "c", "./deployment/local/config.yaml", "configuration file path")
	rootCmd.AddCommand(apiCmd, boAPICmd)
}
