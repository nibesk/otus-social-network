package app

import (
	"log"
	"net/http"
	"service-chat/app/config"
	"service-chat/app/web"
	"service-chat/app/webSocket"
)

type App struct {
	dispatcher   web.Dispatcher
	webSocketHub *webSocket.Hub
}

func Init() *App {
	config.InitConfig()

	return &App{
		dispatcher:   web.InitDispatcher(),
		webSocketHub: webSocket.InitHub(),
	}
}

func (a *App) Run() {
	go a.webSocketHub.Run()
	log.Println("web-socket service has been started")

	log.Fatal(http.ListenAndServe(":"+config.Env.Server.Port, a.dispatcher.Router))
}
