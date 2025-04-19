package service

import (
	"context"
	"log"
	"sync"

	"github.com/MomsEngineer/scheduler_service/api_gateway/config"
	"github.com/MomsEngineer/scheduler_service/api_gateway/models"
	"github.com/MomsEngineer/scheduler_service/api_gateway/utils"
	"github.com/MomsEngineer/scheduler_service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AppointmentCreator interface {
	CreateAppointment(req *models.AppointmentRequest) (*models.AppointmentResponse, error)
}

type GRPCAppointmentService struct{}

var (
	client  proto.AppointmentServiceClient
	once    sync.Once
	initErr error
)

func initGRPCClient() {
	conn, err := grpc.NewClient(config.AppConfig.AppointmentServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		initErr = err
		return
	}
	client = proto.NewAppointmentServiceClient(conn)
}

func (s *GRPCAppointmentService) CreateAppointment(req *models.AppointmentRequest) (*models.AppointmentResponse, error) {
	once.Do(initGRPCClient)
	if initErr != nil {
		log.Println("gRPC client not initialized:", initErr)
		return nil, initErr
	}

	protoReq, err := utils.AppointmentRequestToProto(req)
	if err != nil {
		log.Println("Faild to convert a request to proto", err)
		return nil, err
	}

	res, err := client.CreateAppointment(context.TODO(), protoReq)
	if err != nil {
		log.Println("Faild to create a new appointment", err)
		return nil, err
	}

	return utils.ProtoToAppointmentResponse(res)
}
