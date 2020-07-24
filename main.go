package main

import (
	"github.com/badThug/otus-social-network/app"
	config2 "github.com/badThug/otus-social-network/app/components/config"
)

func main() {
	config := config2.InitConfig()

	app := app.Init(config)
	app.Run()
}
