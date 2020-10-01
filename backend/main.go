package main

import (
	"github.com/badThug/otus-social-network/app"
	"github.com/badThug/otus-social-network/app/config"
	"log"
)

func main() {
	log.Println("Server has been started")

	config := config.InitConfig()
	app := app.Init(config)
	app.Run()
}
