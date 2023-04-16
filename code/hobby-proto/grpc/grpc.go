package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"gitlab.com/book6485215/user_go_user_service/config"
	"gitlab.com/book6485215/user_go_user_service/genproto/user_service"
	"gitlab.com/book6485215/user_go_user_service/grpc/client"
	"gitlab.com/book6485215/user_go_user_service/grpc/service"
	"gitlab.com/book6485215/user_go_user_service/pkg/logger"
	"gitlab.com/book6485215/user_go_user_service/storage"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	user_service.RegisterUserServiceServer(grpcServer, service.NewUserService(cfg, log, strg, srvc))
	user_service.RegisterHobbyServiceServer(grpcServer, service.NewHobbyService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
