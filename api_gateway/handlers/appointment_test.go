package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MomsEngineer/scheduler_service/api_gateway/handlers"
	"github.com/MomsEngineer/scheduler_service/api_gateway/mock"
	"github.com/MomsEngineer/scheduler_service/api_gateway/models"
	"github.com/stretchr/testify/assert"
)

func TestAppointments(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		endpoint       string
		body           []byte
		expectedStatus int
		expectedBody   *models.AppointmentResponse
		expectedError  *models.APIError
	}{
		{
			name:           "GET request",
			method:         http.MethodGet,
			endpoint:       "/appointments",
			body:           nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   nil,
			expectedError: &models.APIError{
				Code:    "method_not_allowed",
				Message: "Only POST requests are allowed",
			},
		},
		{
			name:           "POST request with empty body",
			method:         http.MethodPost,
			endpoint:       "/appointments",
			body:           nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
			expectedError: &models.APIError{
				Code:    "invalid_json",
				Message: "Failed to parse request body",
			},
		},
		{
			name:           "POST request with wrong body",
			method:         http.MethodPost,
			endpoint:       "/appointments",
			body:           []byte(`{bad json}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
			expectedError: &models.APIError{
				Code:    "invalid_json",
				Message: "Failed to parse request body",
			},
		},
		{
			name:           "POST valid JSON",
			method:         http.MethodPost,
			endpoint:       "/appointments",
			body:           []byte(`{"user_id":"123","doctor_id":"456","datetime":"2025-04-12T10:00:00Z"}`),
			expectedStatus: http.StatusCreated,
			expectedBody: &models.AppointmentResponse{
				AppointmentID: "789",
				Status:        "Created",
			},
			expectedError: nil,
		},
		{
			name:           "POST with missing UserID",
			method:         http.MethodPost,
			endpoint:       "/appointments",
			body:           []byte(`{"doctor_id":"456","datetime":"2025-04-12T10:00:00Z"}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
			expectedError: &models.APIError{
				Code:    "validation_error",
				Message: "Missing required fields",
			},
		},
		{
			name:           "POST with missing DoctorID",
			method:         http.MethodPost,
			endpoint:       "/appointments",
			body:           []byte(`{"user_id":"123","datetime":"2025-04-12T10:00:00Z"}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
			expectedError: &models.APIError{
				Code:    "validation_error",
				Message: "Missing required fields",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.endpoint, bytes.NewBuffer(tt.body))
			res := httptest.NewRecorder()
			handlers.CreateAppointmentHandler(res, req, &mock.MockService{})

			assert.Equal(t, tt.expectedStatus, res.Code)

			var respBody models.APIResponse[models.AppointmentResponse]
			err := json.Unmarshal(res.Body.Bytes(), &respBody)
			assert.NoError(t, err)
			if tt.expectedBody != nil {
				assert.Equal(t, "Created", respBody.Data.Status)
				assert.NotEmpty(t, respBody.Data.AppointmentID)
			} else if tt.expectedError != nil {
				assert.Equal(t, tt.expectedError, respBody.Error)
			}
		})
	}
}
