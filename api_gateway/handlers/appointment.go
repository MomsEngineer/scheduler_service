package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MomsEngineer/scheduler_service/api_gateway/models"
	"github.com/go-playground/validator/v10"
)

func CreateAppointmentHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
		log.Println("CreateAppointmentHandler(): Error: Wrong HTTP method")
		return
	}

	request := models.AppointmentRequest{}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Failed to parse JSON", http.StatusBadRequest)
		log.Println("CreateAppointmentHandler(): Error: Failed to parse JSON")
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	response := models.AppointmentResponse{
		AppointmentResponseData: models.AppointmentResponseData{
			AppointmentID: "789",
			Status:        "Created",
		},
		Error: nil,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
		log.Println("CreateAppointmentHandler(): Error: Failed to encode JSON")
		return
	}
}
