package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Product() ProductRepoI
	Category() CategoryRepoI
	Brand() BrandRepoI
	Stock() StockRepoI
	Store() StoreRepoI
	Customer() CustomerRepoI
	Staff() StaffRepoI
	Order() OrderRepoI
	PromoCode() PromoCodeRepoI
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (int, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error)
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (int, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Delete(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateCategory) (int64, error)
}

type BrandRepoI interface {
	Create(context.Context, *models.CreateBrand) (int, error)
	GetByID(context.Context, *models.BrandPrimaryKey) (*models.Brand, error)
	GetList(context.Context, *models.GetListBrandRequest) (*models.GetListBrandResponse, error)
	Update(ctx context.Context, req *models.UpdateBrand) (int64, error)
	Delete(ctx context.Context, req *models.BrandPrimaryKey) (int64, error)
}

type StockRepoI interface {
	Create(ctx context.Context, req *models.CreateStock) (int, int, error)
	GetByID(ctx context.Context, req *models.StockPrimaryKey) (*models.GetStock, error)
	GetList(ctx context.Context, req *models.GetListStockRequest) (resp *models.GetListStockResponse, err error)
	Update(ctx context.Context, req *models.UpdateStock) (int64, error)
	Delete(ctx context.Context, req *models.StockPrimaryKey) (int64, error)
	SendProduct(ctx context.Context, req *models.SendProduct) (int64, error)
	GetProductsFormStock(ctx context.Context, req *models.CreateOrderItem) error
}

type StoreRepoI interface {
	Create(ctx context.Context, req *models.CreateStore) (int, error)
	GetByID(ctx context.Context, req *models.StorePrimaryKey) (*models.Store, error)
	GetList(ctx context.Context, req *models.GetListStoreRequest) (resp *models.GetListStoreResponse, err error)
	UpdatePut(ctx context.Context, req *models.UpdateStore) (int64, error)
	UpdatePatch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(ctx context.Context, req *models.StorePrimaryKey) (int64, error)
}

type CustomerRepoI interface {
	Create(ctx context.Context, req *models.CreateCustomer) (int, error)
	GetByID(ctx context.Context, req *models.CustomerPrimaryKey) (*models.Customer, error)
	GetList(ctx context.Context, req *models.GetListCustomerRequest) (resp *models.GetListCustomerResponse, err error)
	UpdatePut(ctx context.Context, req *models.UpdateCustomer) (int64, error)
	UpdatePatch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(ctx context.Context, req *models.CustomerPrimaryKey) (int64, error)
}

type StaffRepoI interface {
	Create(ctx context.Context, req *models.CreateStaff) (int, error)
	GetByID(ctx context.Context, req *models.StaffPrimaryKey) (*models.Staff, error)
	GetList(ctx context.Context, req *models.GetListStaffRequest) (resp *models.GetListStaffResponse, err error)
	UpdatePut(ctx context.Context, req *models.UpdateStaff) (int64, error)
	UpdatePatch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(ctx context.Context, req *models.StaffPrimaryKey) (int64, error)
}

type OrderRepoI interface {
	Create(ctx context.Context, req *models.CreateOrder) (int, error)
	GetByID(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error)
	GetList(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error)
	Update(ctx context.Context, req *models.UpdateOrder) (int64, error)
	UpdatePatch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(ctx context.Context, req *models.OrderPrimaryKey) (int64, error)
	AddOrderItem(ctx context.Context, req *models.CreateOrderItem) error
	RemoveOrderItem(ctx context.Context, req *models.OrderItemPrimaryKey) error
	InfoOfSoldProductsByStaffer(ctx context.Context, req *models.GetListEmployeeRequest) (resp *models.GetListEmployeeResponse, err error)
	TotalPriceWithOrder(ctx context.Context, req *models.OrderTotalPrice) (string, error)
}
type PromoCodeRepoI interface {
	CreatePromoCode(ctx context.Context, req *models.CreatePromoCode) (string, error)
	GetByIDPromoCode(ctx context.Context, req *models.PromoCodePrimaryKey) (*models.PromoCode, error)
	GetListPromoCode(ctx context.Context, req *models.GetListPromoCodeRequest) (resp *models.GetListPromoCodeResponse, err error)
	DeletePromoCode(ctx context.Context, req *models.PromoCodePrimaryKey) (int64, error)
}