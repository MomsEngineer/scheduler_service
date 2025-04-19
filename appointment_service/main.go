package main

import (
	"log"
	"net"

	"github.com/MomsEngineer/scheduler_service/appointment_service/proto"
	"github.com/MomsEngineer/scheduler_service/appointment_service/server"
	"google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()

	proto.RegisterAppointmentServiceServer(grpcServer, &server.Server{})

	log.Println("gRPC server listening on :50051")
	grpcServer.Serve(lis)
}
