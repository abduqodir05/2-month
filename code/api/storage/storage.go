package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
	User() UserRepoI
	Category() CategoryRepoI
	Courier() CourierRepoI
	Customer() CustomerRepoI
	Product() ProductRepoI
	Order() OrderRepoI
}

//!! Book-------------------------------
type BookRepoI interface {
	Create(context.Context, *models.CreateBook) (string, error)
	GetByID(context.Context, *models.BookPrimaryKey) (*models.Book, error)
	GetList(context.Context, *models.GetListBookRequest) (*models.GetListBookResponse, error)
	Update(context.Context, *models.UpdateBook) (int64, error)
	Patch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.BookPrimaryKey) error
}

//!! User--------------------------------
type UserRepoI interface {
	CreateUser(ctx context.Context, req *models.CreateUser) (string, error)
	GetByIdUser(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetListUser(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) 
	UpdateUser(ctx context.Context, req *models.UpdateUser) (int64, error)
	PatchUser(ctx context.Context, req *models.PatchRequest) (int64, error)
	DeleteUser(ctx context.Context, req *models.UserPrimaryKey) error
}

//!! Category--------------------------------
type CategoryRepoI interface {
	CreateCategory(ctx context.Context, req *models.CreateCategory) (string, error)
	GetByIdCategory(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error)
	GetListCategory(ctx context.Context, req *models.GetListCategoryRequest) (resp *models.GetListCategoryResponse, err error) 
	UpdateCategory(ctx context.Context, req *models.UpdateCategory) (int64, error)
	PatchCategory(ctx context.Context, req *models.PatchRequest) (int64, error)
	DeleteCategory(ctx context.Context, req *models.CategoryPrimaryKey) error
}

//!! Courier--------------------------------
type CourierRepoI interface {
	CreateCourier(ctx context.Context, req *models.CreateCourier) (string, error)
	GetByIdCourier(ctx context.Context, req *models.CourierPrimaryKey) (*models.Courier, error)
	GetListCourier(ctx context.Context, req *models.GetListCourierRequest) (resp *models.GetListCourierResponse, err error) 
	UpdateCourier(ctx context.Context, req *models.UpdateCourier) (int64, error)
	PatchCourier(ctx context.Context, req *models.PatchRequest) (int64, error)
	DeleteCourier(ctx context.Context, req *models.CourierPrimaryKey) error
}

//!! Customer--------------------------------
type CustomerRepoI interface {
	CreateCustomer(ctx context.Context, req *models.CreateCustomer) (string, error)
	GetByIdCustomer(ctx context.Context, req *models.CustomerPrimaryKey) (*models.Customer, error)
	GetListCustomer(ctx context.Context, req *models.GetListCustomerRequest) (resp *models.GetListCustomerResponse, err error) 
	UpdateCustomer(ctx context.Context, req *models.UpdateCustomer) (int64, error)
	PatchCustomer(ctx context.Context, req *models.PatchRequest) (int64, error)
	DeleteCustomer(ctx context.Context, req *models.CustomerPrimaryKey) error
}

//!! Product--------------------------------
type ProductRepoI interface {
	CreateProduct(ctx context.Context, req *models.CreateProduct) (string, error)
	GetByIdProduct(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error)
	GetListProduct(ctx context.Context, req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error) 
	UpdateProduct(ctx context.Context, req *models.UpdateProduct) (int64, error)
	PatchProduct(ctx context.Context, req *models.PatchRequest) (int64, error)
	DeleteProduct(ctx context.Context, req *models.ProductPrimaryKey) error
}

//!! Order--------------------------------
type OrderRepoI interface {
	CreateOrder(ctx context.Context, req *models.CreateOrder) (string, error)
	GetByIdOrder(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error)
	GetListOrder(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error) 
	UpdateOrder(ctx context.Context, req *models.UpdateOrder) (int64, error)
	PatchOrder(ctx context.Context, req *models.PatchRequest) (int64, error)
	DeleteOrder(ctx context.Context, req *models.OrderPrimaryKey) error
}
