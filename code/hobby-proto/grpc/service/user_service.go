package service

import (
	"context"

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

type UserService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *UserService {
	return &UserService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *UserService) Create(ctx context.Context, req *user_service.CreateUser) (resp *user_service.User, err error) {

	i.log.Info("---CreateUser------>", logger.Any("req", req))

	pKey, err := i.strg.User().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateUser->User->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.User().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *UserService) GetByID(ctx context.Context, req *user_service.UserPrimaryKey) (resp *user_service.User, err error) {

	i.log.Info("---GetUserByID------>", logger.Any("req", req))

	resp, err = i.strg.User().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetUserByID->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *UserService) GetList(ctx context.Context, req *user_service.GetListUserRequest) (resp *user_service.GetListUserResponse, err error) {

	i.log.Info("---GetUsers------>", logger.Any("req", req))

	resp, err = i.strg.User().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetUsers->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *UserService) Update(ctx context.Context, req *user_service.UpdateUser) (resp *user_service.User, err error) {

	i.log.Info("---UpdateUser------>", logger.Any("req", req))

	rowsAffected, err := i.strg.User().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.User().GetByPKey(ctx, &user_service.UserPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *UserService) UpdatePatch(ctx context.Context, req *user_service.UpdatePatchUser) (resp *user_service.User, err error) {

	i.log.Info("---UpdatePatchEnrolledStudent------>", logger.Any("req", req))

	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.User().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.User().GetByPKey(ctx, &user_service.UserPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *UserService) Delete(ctx context.Context, req *user_service.UserPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteUser------>", logger.Any("req", req))

	err = i.strg.User().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
