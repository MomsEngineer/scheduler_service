package mock

import "github.com/MomsEngineer/scheduler_service/api_gateway/models"

type MockService struct{}

func (s *MockService) CreateAppointment(_ *models.AppointmentRequest) (*models.AppointmentResponse, error) {
	return &models.AppointmentResponse{
		AppointmentID: "abc123",
		Status:        "Created",
	}, nil
}
