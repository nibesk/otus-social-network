package app

import (
	config2 "github.com/badThug/otus-social-network/app/components/config"
	"github.com/badThug/otus-social-network/app/components/storage"
	"github.com/badThug/otus-social-network/app/components/web"
	"log"
	"net/http"
)

type App struct {
	config     *config2.Config
	db         *storage.Connection
	dispatcher web.Dispatcher
}

func Init(config *config2.Config) *App {
	db := storage.ConnectDatabase(config.DB)
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
