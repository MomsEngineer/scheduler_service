package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MomsEngineer/scheduler_service/api_gateway/handlers"
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
		expectedBody   string
	}{
		{
			name:           "GET request",
			method:         http.MethodGet,
			endpoint:       "/appointments",
			body:           nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Only POST requests are allowed!",
		},
		{
			name:           "POST request with empty body",
			method:         http.MethodPost,
			endpoint:       "/appointments",
			body:           nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Failed to parse JSON",
		},
		{
			name:           "POST request with wrong body",
			method:         http.MethodPost,
			endpoint:       "/appointments",
			body:           []byte(`{bad json}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Failed to parse JSON",
		},
		{
			name:           "POST valid JSON",
			method:         http.MethodPost,
			endpoint:       "/appointments",
			body:           []byte(`{"user_id":"123","doctor_id":"456","datetime":"2025-04-12T10:00:00Z"}`),
			expectedStatus: http.StatusOK,
			expectedBody:   `"status":"Created"`,
		},
		{
			name:           "POST with missing UserID",
			method:         http.MethodPost,
			endpoint:       "/appointments",
			body:           []byte(`{"doctor_id":"456","datetime":"2025-04-12T10:00:00Z"}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request",
		},
		{
			name:           "POST with missing DoctorID",
			method:         http.MethodPost,
			endpoint:       "/appointments",
			body:           []byte(`{"user_id":"123","datetime":"2025-04-12T10:00:00Z"}`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.endpoint, bytes.NewBuffer(tt.body))
			res := httptest.NewRecorder()
			handlers.CreateAppointmentHandler(res, req)

			assert.Equal(t, tt.expectedStatus, res.Code)
			if tt.expectedStatus == http.StatusOK {
				var respBody models.AppointmentResponse
				err := json.Unmarshal(res.Body.Bytes(), &respBody)
				assert.NoError(t, err)
				assert.Equal(t, "Created", respBody.Status)
				assert.NotEmpty(t, respBody.AppointmentID)
			} else if tt.expectedBody != "" {
				assert.Contains(t, res.Body.String(), tt.expectedBody)
			}
		})
	}
}
