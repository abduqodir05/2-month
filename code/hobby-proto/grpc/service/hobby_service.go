package service

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/book6485215/user_go_user_service/config"
	"gitlab.com/book6485215/user_go_user_service/genproto/user_service"
	"gitlab.com/book6485215/user_go_user_service/grpc/client"
	"gitlab.com/book6485215/user_go_user_service/models"
	"gitlab.com/book6485215/user_go_user_service/pkg/logger"
	"gitlab.com/book6485215/user_go_user_service/storage"
)

type HobbyService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedHobbyServiceServer
}

func NewHobbyService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *HobbyService {
	// fmt.Println("hobby service intitialization")
	return &HobbyService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *HobbyService) CreateHobby(ctx context.Context, req *user_service.CreateHobby) (resp *user_service.Hobby, err error) {
	fmt.Println("this is println")
	i.log.Info("---CreateHobby------>", logger.Any("req", req))

	pKey, err := i.strg.Hobby().CreateHobby(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateHobby->Hobby->CreateHobby--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Hobby().GetByPKeyHobby(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyHobby->Hobby->GetHobby--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *HobbyService) GetByIDHobby(ctx context.Context, req *user_service.HobbyPrimaryKey) (resp *user_service.Hobby, err error) {

	i.log.Info("---GetHobbyByID------>", logger.Any("req", req))

	resp, err = i.strg.Hobby().GetByPKeyHobby(ctx, req)
	if err != nil {
		i.log.Error("!!!GetHobbyByID->Hobby->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *HobbyService) GetListHobby(ctx context.Context, req *user_service.GetListHobbyRequest) (resp *user_service.GetListHobbyResponse, err error) {

	i.log.Info("---GetUsers------>", logger.Any("req", req))

	resp, err = i.strg.Hobby().GetAllHobby(ctx, req)
	if err != nil {
		i.log.Error("!!!GetHobby->Hobby->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *HobbyService) UpdateHobby(ctx context.Context, req *user_service.UpdateHobby) (resp *user_service.Hobby, err error) {

	i.log.Info("---UpdateHobby------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Hobby().UpdateHobby(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateHobby--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Hobby().GetByPKeyHobby(ctx, &user_service.HobbyPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *HobbyService) UpdatePatchHobby(ctx context.Context, req *user_service.UpdatePatchHobby) (resp *user_service.Hobby, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))

	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Hobby().UpdatePatchHobby(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Hobby().GetByPKeyHobby(ctx, &user_service.HobbyPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetHobby->Hobby->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *HobbyService) DeleteHobby(ctx context.Context, req *user_service.HobbyPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteUser------>", logger.Any("req", req))

	err = i.strg.Hobby().DeleteHobby(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteHobby->Hobby->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
