package app

import (
	"github.com/gorilla/mux"
	"github.com/soarex16/fabackend/routes"
)

type App struct {
	Config *Config
	Router *mux.Router
}

func New() (app *App, err error) {
	app = &App{}
	app.Config, err = InitConfiguration()

	if err != nil {
		return nil, err
	}

	//TODO: коннект в субд, и т.д.

	app.Router = routes.New()

	return app, err
}

//TODO: метод, который осуществляет корректное завершение работы приложения
