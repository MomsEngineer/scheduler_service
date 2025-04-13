package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MomsEngineer/scheduler_service/api_gateway/handlers"
	"github.com/stretchr/testify/assert"
)

func TestHealthzHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	res := httptest.NewRecorder()

	handlers.Healthz(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
