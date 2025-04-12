package main

import (
	"net/http"

	"log"

	handler "github.com/MomsEngineer/scheduler_service/api_gateway/handlers"
)

func main() {
	log.Println("main: Start API-Gateway service")
	mux := http.NewServeMux()

	mux.HandleFunc(`/healthz`, handler.Healthz)
	mux.HandleFunc(`/appointments`, handler.CreateAppointmentHandler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}

	log.Println("main: Finish API-Gateway service")
}
