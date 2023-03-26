package storage

import "app/api/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	CreateUser(*models.CreateUser) (string, error)
	GetByIDUser(*models.UserPrimaryKey) (*models.User, error)
	UpdateUser(*models.UpdateUser) (string, error)
	DeleteUser(*models.DeleteUser) error
	GetListUser(*models.GetListUserRequest) (*models.GetListUserResponse, error)
}