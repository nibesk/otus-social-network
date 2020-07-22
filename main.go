package main

import (
	"github.com/badThug/otus-social-network/app"
	"github.com/badThug/otus-social-network/config"
)

//import (
//	"net/http"
//	"os"
//)
//
//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("<h1>Hello World!</h1>"))
//}
//
//func main() {
//	port := os.Getenv("PORT")
//	if port == "" {
//		port = "3000"
//	}
//
//	mux := http.NewServeMux()
//
//	mux.HandleFunc("/", indexHandler)
//	http.ListenAndServe(":"+port, mux)
//}

func main() {
	config := config.Init()

	app := app.Init(config)
	app.Run()
}
