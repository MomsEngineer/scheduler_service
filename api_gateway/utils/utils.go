package utils

import (
	"bytes"
	"encoding/json"
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

func ToProto(r *models.AppointmentRequest) *proto.CreateAppointmentRequest {
	return &proto.CreateAppointmentRequest{
		UserId:   r.UserID,
		DoctorId: r.DoctorID,
		DateTime: r.DateTime,
	}
}

func FromProto(res *proto.CreateAppointmentResponse) *models.AppointmentResponse {
	return &models.AppointmentResponse{
		AppointmentID: res.AppointmentId,
		Status:        res.Status,
	}
}
