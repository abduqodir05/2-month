package storage

import (
	"context"

	"gitlab.com/book6485215/user_go_user_service/models"

	"gitlab.com/book6485215/user_go_user_service/genproto/user_service"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Hobby() HobbyRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *user_service.CreateUser) (resp *user_service.UserPrimaryKey, err error)
	GetByPKey(ctx context.Context, req *user_service.UserPrimaryKey) (resp *user_service.User, err error)
	GetAll(ctx context.Context, req *user_service.GetListUserRequest) (resp *user_service.GetListUserResponse, err error)
	Update(ctx context.Context, req *user_service.UpdateUser) (rowsAffected int64, err error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *user_service.UserPrimaryKey) error
}
type HobbyRepoI interface {
	CreateHobby(ctx context.Context, req *user_service.CreateHobby) (resp *user_service.HobbyPrimaryKey, err error)
	GetByPKeyHobby(ctx context.Context, req *user_service.HobbyPrimaryKey) (resp *user_service.Hobby, err error)
	GetAllHobby(ctx context.Context, req *user_service.GetListHobbyRequest) (resp *user_service.GetListHobbyResponse, err error)
	UpdateHobby(ctx context.Context, req *user_service.UpdateHobby) (rowsAffected int64, err error) 
	UpdatePatchHobby(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error)
	DeleteHobby(ctx context.Context, req *user_service.HobbyPrimaryKey) error 
}
