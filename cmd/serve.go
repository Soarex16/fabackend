package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/app"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the web api",
	RunE: func(cmd *cobra.Command, args []string) error {
		application, err := app.New()

		defer func() {
			logrus.Info("Closing database connection")
			application.Context.DB.Close()
		}()

		if err != nil {
			logrus.Fatalf("Unable to start application: %v", err)
		}

		srv := newServer(application)
		// start server in it's own goroutine
		go func() {
			if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logrus.Fatal(err)
			}
		}()
		logrus.Infof("Server has been started ad port %v", application.Config.Port)

		// subscribe to interrupt
		osSig := make(chan os.Signal, 1)
		signal.Notify(osSig, os.Interrupt)

		<-osSig
		logrus.Info("Caught Interrupt signal. Shutting down")

		// context for gracefull shutdown
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if err = srv.Shutdown(ctx); err != nil {
			err = nil
			logrus.Errorf("Server unsuccessfully shutted down with error: %v", err)
		}

		return err
	},
}

func newServer(app *app.App) *http.Server {
	addr := fmt.Sprintf(":%v", (*app).Config.Port)
	server := http.Server{
		Addr:    addr,
		Handler: (*app).Router,
	}

	return &server
}

//TODO: метод, который осуществляет корректное завершение работы приложения (и закрывает коннект к бд)

func init() {
	rootCmd.AddCommand(serveCmd)
}
