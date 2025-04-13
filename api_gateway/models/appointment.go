package models

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type APIResponse[T any] struct {
	Data  *T        `json:"data,omitempty"`
	Error *APIError `json:"error,omitempty"`
}

type AppointmentRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	DoctorID string `json:"doctor_id" validate:"required"`
	DateTime string `json:"datetime"`
}

type AppointmentResponse struct {
	AppointmentID string `json:"appointment_id"`
	Status        string `json:"status"`
}
