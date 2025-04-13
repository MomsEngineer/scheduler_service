package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MomsEngineer/scheduler_service/api_gateway/utils"
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
