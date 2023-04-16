package client

import (
	"gitlab.com/book6485215/user_go_user_service/config"
)

type ServiceManagerI interface{}

type grpcClients struct{}

	func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

		return &grpcClients{}, nil
	}
