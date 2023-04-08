package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI

}

// type ProductRepoI interface {
// 	Create(context.Context, *models.CreateProduct) (int, error)
// 	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
// 	GetList(context.Context, *models.GetListProductRequest) (*models.GetListProductResponse, error)
// 	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
// 	Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error)
// }

type UserRepoI interface{
	CreateUser(ctx context.Context, req *models.CreateUser) (string, error)
	GetByIDUser(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetListUser(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error)
}

