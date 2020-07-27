package app

import (
	"github.com/badThug/otus-social-network/app/config"
	"github.com/badThug/otus-social-network/app/storage"
	"github.com/badThug/otus-social-network/app/web"
	"log"
	"net/http"
)

type App struct {
	config     *config.Config
	db         *storage.DbConnection
	dispatcher web.Dispatcher
}

func Init(config *config.Config) *App {
	db := storage.CreateDbConnection(config.DB)
	dispatcher := web.InitDispatcher(db, config)

	app := &App{
		db:         db,
		config:     config,
		dispatcher: dispatcher,
	}

	return app
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":"+a.config.Server.HttpPort, a.dispatcher.Router))
}
