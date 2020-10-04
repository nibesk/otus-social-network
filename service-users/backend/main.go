package main

import (
	"log"
	"service-users/app"
)

func main() {
	log.Println("service-users has been started")

	app := app.Init()
	app.Run()
}
