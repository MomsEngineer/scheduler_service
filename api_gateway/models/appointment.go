package models

type AppointmentRequest struct {
	UserID   string `json:"user_id"`
	DoctorID string `json:"doctor_id"`
	DateTime string `json:"datetime"`
}

type AppointmentResponseData struct {
	Appointment string `json:"appointment_id"`
	Status      string `json:"status"`
}

type AppointmentResponse struct {
	AppointmentResponseData
	Error error `json:"error,omitempty"`
}
