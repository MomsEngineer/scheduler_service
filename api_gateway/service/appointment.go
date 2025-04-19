package service

import (
	"context"
	"log"
	"sync"

	"github.com/MomsEngineer/scheduler_service/api_gateway/models"
	"github.com/MomsEngineer/scheduler_service/api_gateway/proto"
	"github.com/MomsEngineer/scheduler_service/api_gateway/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client  proto.AppointmentServiceClient
	once    sync.Once
	initErr error
)

func initGRPCClient() {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		initErr = err
		return
	}
	client = proto.NewAppointmentServiceClient(conn)
}

func CreateAppointment(req *models.AppointmentRequest) (*models.AppointmentResponse, error) {
	once.Do(initGRPCClient)
	if initErr != nil {
		log.Println("gRPC client not initialized:", initErr)
		return nil, initErr
	}

	res, err := client.CreateAppointment(context.TODO(), utils.ToProto(req))
	if err != nil {
		log.Println("Faild to create a new appointment", err)
		return nil, err
	}

	return utils.FromProto(res), nil
}
