package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the web api",
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("Starting server")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
