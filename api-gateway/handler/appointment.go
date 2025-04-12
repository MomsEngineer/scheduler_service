package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	User     string `json:"user_id"`
	Doctor   string `json:"doctor_id"`
	DateTime string `json:"datetime"`
}

type ResponseData struct {
	Appointment string `json:"appointment_id"`
	Status      string `json:"status"`
}

type Response struct {
	ResponseData
	Error error `json:"error,omitempty"`
}

func Appointments(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
		log.Println("Appointments(): Error: Wrong HTTP method")
		return
	}

	request := Request{}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Failed to parse JSON", http.StatusBadRequest)
		log.Println("Appointments(): Error: Failed to parse JSON")
		return
	}

	response := Response{
		ResponseData: ResponseData{
			Appointment: "789",
			Status:      "Created",
		},
		Error: nil,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
		log.Println("Appointments(): Error: Failed to encode JSON")
		return
	}
}
