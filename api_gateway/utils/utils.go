package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/MomsEngineer/scheduler_service/api_gateway/models"
	"github.com/MomsEngineer/scheduler_service/api_gateway/proto"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		log.Printf("failed to encode JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

func AppointmentRequestToProto(req *models.AppointmentRequest) (*proto.CreateAppointmentRequest, error) {
	if req == nil {
		log.Println("ERROR: AppointmentRequestToProto(): request is NIL")
		return nil, errors.New("request is nil")
	}
	return &proto.CreateAppointmentRequest{
		UserId:   req.UserID,
		DoctorId: req.DoctorID,
		DateTime: req.DateTime,
	}, nil
}

func ProtoToAppointmentResponse(res *proto.CreateAppointmentResponse) (*models.AppointmentResponse, error) {
	if res == nil {
		log.Println("ERROR: ProtoToAppointmentResponse(): response is NIL")
		return nil, errors.New("response is nil")
	}
	return &models.AppointmentResponse{
		AppointmentID: res.AppointmentId,
		Status:        res.Status,
	}, nil
}
