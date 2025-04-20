package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/MomsEngineer/scheduler_service/api_gateway/models"
	"github.com/stretchr/testify/require"
)

func TestCreateAppointment_E2E(t *testing.T) {
	addr := os.Getenv("API_GATEAY_ADDR")
	if addr == "" {
		addr = "http://localhost:8080"
	}

	payload := []byte(`{"user_id":"123","doctor_id":"456","datetime":"2025-04-12T10:00:00Z"}`)

	res, err := http.Post(addr+"/appointments", "application/json", bytes.NewBuffer(payload))
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusCreated, res.StatusCode)

	var resp models.APIResponse[models.AppointmentResponse]
	err = json.NewDecoder(res.Body).Decode(&resp)
	require.NoError(t, err)
	require.Equal(t, "created", resp.Data.Status)
	require.NotEmpty(t, resp.Data.AppointmentID)
}
