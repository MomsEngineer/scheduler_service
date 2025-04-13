package models

type AppointmentRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	DoctorID string `json:"doctor_id" validate:"required"`
	DateTime string `json:"datetime"`
}

type AppointmentResponseData struct {
	AppointmentID string `json:"appointment_id"`
	Status        string `json:"status"`
}

type AppointmentResponse struct {
	AppointmentResponseData
	Error error `json:"error,omitempty"`
}
