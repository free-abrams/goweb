package main

import (
	"goweb/routes"
	"log"
	"net/http"
)

func main() {

	startWebServer("9090")
}

func startWebServer(port string) {
	r := routes.NewRouter()

	http.Handle("/", r)
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
