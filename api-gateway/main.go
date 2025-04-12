package main

import (
	"api-gateway/handler"
	"net/http"

	"log"
)

func main() {
	log.Println("main: Start API-Gateway service")
	mux := http.NewServeMux()

	mux.HandleFunc(`/healthz`, handler.Health)
	mux.HandleFunc(`/appointments`, handler.Appointments)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}

	log.Println("main: Finish API-Gateway service")
}
