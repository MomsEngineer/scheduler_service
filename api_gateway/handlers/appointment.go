package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MomsEngineer/scheduler_service/api_gateway/models"
	"github.com/MomsEngineer/scheduler_service/api_gateway/utils"
	"github.com/go-playground/validator/v10"
)

func CreateAppointmentHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		utils.WriteJSON(res, http.StatusMethodNotAllowed, models.APIResponse[any]{
			Error: &models.APIError{
				Code:    "method_not_allowed",
				Message: "Only POST requests are allowed",
			},
		})
		log.Println("CreateAppointmentHandler(): Error: Wrong HTTP method")
		return
	}

	request := models.AppointmentRequest{}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		utils.WriteJSON(res, http.StatusBadRequest, models.APIResponse[any]{
			Error: &models.APIError{
				Code:    "invalid_json",
				Message: "Failed to parse request body",
			},
		})
		log.Println("CreateAppointmentHandler(): Error: Failed to parse JSON")
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		utils.WriteJSON(res, http.StatusBadRequest, models.APIResponse[any]{
			Error: &models.APIError{
				Code:    "validation_error",
				Message: "Missing required fields",
			},
		})
		return
	}

	response := models.APIResponse[models.AppointmentResponse]{
		Data: &models.AppointmentResponse{
			AppointmentID: "789",
			Status:        "Created",
		},
	}

	utils.WriteJSON(res, http.StatusCreated, response)
}
