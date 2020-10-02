package main

import (
	"log"
	"service-users/app"
	"service-users/app/config"
)

func main() {
	log.Println("Server has been started")

	config := config.InitConfig()
	app := app.Init(config)
	app.Run()
}
