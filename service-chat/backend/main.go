package main

import (
	"log"
	"service-chat/app"
)

func main() {
	log.Println("service-chat has been started")

	app := app.Init()
	app.Run()
}
