package main

import (
	"net/http"

	"log"

	"github.com/MomsEngineer/scheduler_service/api_gateway/config"
	"github.com/MomsEngineer/scheduler_service/api_gateway/handlers"
	"github.com/MomsEngineer/scheduler_service/api_gateway/service"
)

func main() {
	config.Init()

	log.Println("main: Start API-Gateway service")
	mux := http.NewServeMux()

	mux.HandleFunc(`/healthz`, handlers.Healthz)

	svc := &service.GRPCAppointmentService{}
	mux.HandleFunc(`/appointments`, func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateAppointmentHandler(w, r, svc)
	})

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}

	log.Println("main: Finish API-Gateway service")
}
