package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/config"
	"app/storage"
)

type Store struct {
	db   *pgxpool.Pool
	product storage.ProductRepoI
	category storage.CategoryRepoI
	brand storage.BrandRepoI
	stocks storage.StocksRepoI
}

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))
	if err != nil {
		return nil, err
	}

	pgpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:   pgpool,
		product: NewProductRepo(pgpool),
		category: NewCategoryRepo(pgpool),
		brand: NewBrandRepo(pgpool),
		stocks: NewStocksRepo(pgpool),

	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}


func (s *Store) Product() storage.ProductRepoI {

	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *Store) Category() storage.CategoryRepoI {

	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *Store) Brand() storage.BrandRepoI {

	if s.brand == nil {
		s.brand = NewBrandRepo(s.db)
	}

	return s.brand
}

func (s *Store) Stocks() storage.StocksRepoI {

	if s.stocks == nil {
		s.stocks = NewStocksRepo(s.db)
	}

	return s.stocks
}

// GORM
// ROW
// SQLBUILDER
// SQLX
// PGXPOOL
