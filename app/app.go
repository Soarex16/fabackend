package app

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	ctx "github.com/soarex16/fabackend/context"
	"github.com/soarex16/fabackend/routes"

	// register postgres driver
	_ "github.com/lib/pq"
)

// App - global application data
type App struct {
	Config  *Config
	Context *ctx.Context
	Router  *mux.Router
}

// New - initializes configuration, routes, db
func New() (app *App, err error) {
	app = &App{}
	app.Config, err = InitConfiguration()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", app.Config.DbConnectionString)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logrus.Fatalf("Cannot ping database")
		return nil, err
	}

	app.Context = &ctx.Context{DB: db}

	appRoutes := routes.GetAll(app.Context)
	app.Router = routes.NewRouter(appRoutes)

	return app, err
}
