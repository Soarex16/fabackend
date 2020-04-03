package app

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"

	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/routes"

	// register postgres driver
	_ "github.com/lib/pq"
)

// App - global application data
type App struct {
	Config *Config
	Store  *Store
	Router *httprouter.Router
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

	app.Store = NewStore(db)

	appRoutes := InitHandlers(app)
	app.Router = routes.NewRouter(appRoutes)

	return app, err
}
