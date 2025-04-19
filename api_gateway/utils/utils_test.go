package utils_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MomsEngineer/scheduler_service/api_gateway/models"
	"github.com/MomsEngineer/scheduler_service/api_gateway/utils"
	"github.com/MomsEngineer/scheduler_service/proto"
	"github.com/stretchr/testify/assert"
)

type broken struct {
	BadField func() // json.Encode не может сериализовать func
}

func TestWriteJSON_Success(t *testing.T) {
	w := httptest.NewRecorder()
	payload := map[string]string{"message": "hello"}

	utils.WriteJSON(w, http.StatusCreated, payload)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"message":"hello"}`, w.Body.String())
}

func TestWriteJSON_EncodingError(t *testing.T) {
	w := httptest.NewRecorder()
	utils.WriteJSON(w, http.StatusOK, broken{})

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Internal Server Error")
}

func TestAppointmentRequestToProto(t *testing.T) {
	tests := []struct {
		name          string
		request       *models.AppointmentRequest
		expectedBody  *proto.CreateAppointmentRequest
		expectedError error
	}{
		{
			name: "Good case",
			request: &models.AppointmentRequest{
				UserID:   "1",
				DoctorID: "2",
				DateTime: "2025-04-12T10:00:00Z",
			},
			expectedBody: &proto.CreateAppointmentRequest{
				UserId:   "1",
				DoctorId: "2",
				DateTime: "2025-04-12T10:00:00Z",
			},
			expectedError: nil,
		},
		{
			name:          "Nil case",
			request:       nil,
			expectedBody:  nil,
			expectedError: errors.New("request is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := utils.AppointmentRequestToProto(tt.request)
			assert.Equal(t, tt.expectedBody, actual)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestProtoToAppointmentResponse(t *testing.T) {
	tests := []struct {
		name          string
		response      *proto.CreateAppointmentResponse
		expectedBody  *models.AppointmentResponse
		expectedError error
	}{
		{
			name: "Good case",
			response: &proto.CreateAppointmentResponse{
				AppointmentId: "1",
				Status:        "created",
			},
			expectedBody: &models.AppointmentResponse{
				AppointmentID: "1",
				Status:        "created",
			},
			expectedError: nil,
		},
		{
			name:          "Nil case",
			response:      nil,
			expectedBody:  nil,
			expectedError: errors.New("response is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := utils.ProtoToAppointmentResponse(tt.response)
			assert.Equal(t, tt.expectedBody, actual)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
