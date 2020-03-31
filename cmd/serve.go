package cmd

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/app"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the web api",
	RunE: func(cmd *cobra.Command, args []string) error {
		application, err := app.New()

		if err != nil {
			logrus.Fatalf("Unable to start application: %v", err)
		}

		defer application.Context.DB.Close()

		addr := fmt.Sprintf(":%v", application.Config.Port)
		logrus.Infof("Starting server at port %v", addr)

		logrus.Fatal(http.ListenAndServe(addr, application.Router))

		return nil
	},
}

//TODO: метод, который осуществляет корректное завершение работы приложения (и закрывает коннект к бд)

func init() {
	rootCmd.AddCommand(serveCmd)
}
