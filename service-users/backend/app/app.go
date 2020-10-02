package app

import (
	"log"
	"net/http"
	"service-users/app/config"
	"service-users/app/storage"
	"service-users/app/web"
)

type App struct {
	config     *config.Config
	db         *storage.DbConnection
	dispatcher web.Dispatcher
}

func Init(config *config.Config) *App {
	db, err := storage.ConnectDb(config.DB)
	if nil != err {
		log.Fatalf("can't connect to database: %+v", err)
	}
	dispatcher := web.InitDispatcher(db, config)

	app := &App{
		db:         db,
		config:     config,
		dispatcher: dispatcher,
	}

	return app
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":3000", a.dispatcher.Router))
}
