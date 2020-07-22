package app

import (
	"github.com/badThug/otus-social-network/app/database"
	"github.com/badThug/otus-social-network/app/dispatcher"
	"github.com/badThug/otus-social-network/config"
	"log"
	"net/http"
)

type App struct {
	config     *config.Config
	db         *database.Connection
	dispatcher *dispatcher.Dispatcher
}

func Init(config *config.Config) *App {
	db := database.Connect(config.DB)
	dispatcher := dispatcher.Init(db)

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
