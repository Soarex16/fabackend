package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "faback",
	Short: "Fitness App backend",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

// Execute - entry point in CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var configFile string

func init() {
	initLogger()
	cobra.OnInitialize(initConfigurationFile)
}

func initConfigurationFile() {
	viper.AutomaticEnv()

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Fatalf("Unable to find configuration file at: %v\n%v", configFile, err)
		} else {
			logrus.Fatalf("Unable to read configuration: %v", err)
		}
	}

}

const logFileDir = "fabackend.log"

func initLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logFile, err := os.OpenFile(logFileDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0644))

	var logWriter io.Writer

	if err != nil {
		logrus.Warningf("Unable to open log file at %v. Using stdout for logging. Err: %v", logFileDir, err)
		logWriter = os.Stdout
	} else {
		logWriter = io.MultiWriter(os.Stdout, logFile)
	}

	logrus.SetOutput(logWriter)
	logrus.SetLevel(logrus.InfoLevel)
}
