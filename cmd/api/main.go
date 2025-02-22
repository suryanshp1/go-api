package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"goapi/internal/handlers"
	log "github.com/sirupsen/logrus"  // alias for logrus is log
)


func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()

	handlers.Handler(r)

	fmt.Println("Starting go rest api server on port 8000..........")

	// generated ascii art for the server starting
	fmt.Println(`
________  ________                 ________  ________  ___     
|\   ____\|\   __  \               |\   __  \|\   __  \|\  \    
\ \  \___|\ \  \|\  \  ____________\ \  \|\  \ \  \|\  \ \  \   
 \ \  \  __\ \  \\\  \|\____________\ \   __  \ \   ____\ \  \  
  \ \  \|\  \ \  \\\  \|____________|\ \  \ \  \ \  \___|\ \  \ 
   \ \_______\ \_______\              \ \__\ \__\ \__\    \ \__\
    \|_______|\|_______|               \|__|\|__|\|__|     \|__|
                                                                `)
	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Fatal(err)
	}
}