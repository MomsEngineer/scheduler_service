package server

import (
	"context"
	"log"

	"github.com/MomsEngineer/scheduler_service/proto"
)

type Server struct {
	proto.UnimplementedAppointmentServiceServer
}

func (s *Server) CreateAppointment(ctx context.Context, req *proto.CreateAppointmentRequest) (*proto.CreateAppointmentResponse, error) {
	log.Println("Creating appointment for:", req.UserId, req.DoctorId)
	return &proto.CreateAppointmentResponse{
		AppointmentId: "abc123",
		Status:        "created",
	}, nil
}
