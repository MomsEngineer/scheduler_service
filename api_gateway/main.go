package main

import (
	"net/http"

	"log"

	handler "github.com/MomsEngineer/scheduler_service/api_gateway/handlers"
	"github.com/MomsEngineer/scheduler_service/api_gateway/service"
)

func main() {
	log.Println("main: Start API-Gateway service")
	mux := http.NewServeMux()

	mux.HandleFunc(`/healthz`, handler.Healthz)

	svc := &service.GRPCAppointmentService{}
	mux.HandleFunc(`/appointments`, func(w http.ResponseWriter, r *http.Request) {
		handler.CreateAppointmentHandler(w, r, svc)
	})

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}

	log.Println("main: Finish API-Gateway service")
}
