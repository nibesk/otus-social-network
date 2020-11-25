package app

import (
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"service-users/app/config"
	"service-users/app/storage"
	"service-users/app/utils"
	"service-users/app/web"
)

type App struct{}

func Init() *App {
	config.InitConfig()
	utils.CreateValidator()

	err := storage.ConnectDb()
	if nil != err {
		log.Fatalf("can't connect to database: %+v", err)
	}

	return &App{}
}

func (a *App) Run() {
	cmdServer := &cobra.Command{
		Use:   "server",
		Short: "Server is a web server for users and feed",
		Run: func(cmd *cobra.Command, args []string) {
			dispatcher := web.InitDispatcher()
			log.Fatal(http.ListenAndServe(":"+config.Env.Server.Port, dispatcher.Router))
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdServer)
	rootCmd.Execute()
}
