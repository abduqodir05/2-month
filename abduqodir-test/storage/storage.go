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
	Stocks() StocksRepoI
}

//!! Product--------------------------------
type ProductRepoI interface {
	CreateProduct(ctx context.Context, req *models.Products) (int, error)
	GetByIdProduct(ctx context.Context, req *models.ProductPrimaryKey) (*models.Products, error)
	GetListProduct(ctx context.Context, req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error)
	UpdateProduct(ctx context.Context, req *models.Products) (int64, error)
	DeleteProduct(ctx context.Context, req *models.ProductPrimaryKey) (int64, error)
}

//!! Category--------------------------------
type CategoryRepoI interface {
	CreateCategory(ctx context.Context, req *models.Category) (int, error)
	GetByIdCategory(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error)
	GetListCategory(ctx context.Context, req *models.GetListCategoryRequest) (resp *models.GetListCategoryResponse, err error)
	UpdateCategory(ctx context.Context, req *models.Category) (int64, error)
	DeleteCategory(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error)
}

//!! Brand--------------------------------
type BrandRepoI interface {
	CreateBrand(ctx context.Context, req *models.Brand) (int, error)
	GetByIdBrand(ctx context.Context, req *models.BrandPrimaryKey) (*models.Brand, error)
	GetListBrand(ctx context.Context, req *models.GetListBrandRequest) (resp *models.GetListBrandResponse, err error)
	UpdateBrand(ctx context.Context, req *models.Brand) (int64, error)
	DeleteBrand(ctx context.Context, req *models.BrandPrimaryKey) (int64, error)
}

//!! Stocks--------------------------------
type StocksRepoI interface {
	CreateStocks(ctx context.Context, req *models.Stock) (int, error)
	GetByIdStocks(ctx context.Context, req *models.StocksPrimaryKey) (*models.Stock, error)
	GetListStocks(ctx context.Context, req *models.GetListStocksRequest) (resp *models.GetListStocksResponse, err error)
	UpdateStocks(ctx context.Context, req *models.Stock) (int64, error)
	DeleteStocks(ctx context.Context, req *models.StocksPrimaryKey) (int64, error)
}
