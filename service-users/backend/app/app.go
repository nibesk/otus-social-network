package app

import (
	"log"
	"net/http"
	"service-users/app/config"
	"service-users/app/storage"
	"service-users/app/utils"
	"service-users/app/web"
)

type App struct {
	db         *storage.DbConnection
	dispatcher web.Dispatcher
}

func Init() *App {
	config.InitConfig()
	utils.CreateValidator()

	db, err := storage.ConnectDb()
	if nil != err {
		log.Fatalf("can't connect to database: %+v", err)
	}
	dispatcher := web.InitDispatcher(db, config.Env)

	app := &App{
		db:         db,
		dispatcher: dispatcher,
	}

	return app
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":"+config.Env.Server.Port, a.dispatcher.Router))
}
