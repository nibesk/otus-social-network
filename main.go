package main

import (
	"github.com/badThug/otus-social-network/app"
	"github.com/badThug/otus-social-network/app/components/config"
)

func main() {
	config := config.InitConfig()

	//test(config)
	//os.Exit(1)

	app := app.Init(config)
	app.Run()

}
