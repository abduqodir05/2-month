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
	book storage.BookRepoI
	user storage.UserRepoI
	category storage.CategoryRepoI
	courier storage.CourierRepoI
	customer storage.CustomerRepoI
	product storage.ProductRepoI
	order storage.OrderRepoI
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
		book: NewBookRepo(pgpool),
		user: NewUserRepo(pgpool),
		category: NewCategoryRepo(pgpool),
		courier: NewCourierRepo(pgpool),
		customer: NewCustomerRepo(pgpool),
		product: NewProductRepo(pgpool),
		order: NewOrderRepo(pgpool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Book() storage.BookRepoI {

	if s.book == nil {
		s.book = NewBookRepo(s.db)
	}

	return s.book
}

func (s *Store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}
func (s *Store) Category() storage.CategoryRepoI {

	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *Store) Courier() storage.CourierRepoI {

	if s.courier == nil {
		s.courier = NewCourierRepo(s.db)
	}

	return s.courier
}

func (s *Store) Customer() storage.CustomerRepoI {

	if s.customer == nil {
		s.customer = NewCustomerRepo(s.db)
	}

	return s.customer
}

func (s *Store) Product() storage.ProductRepoI {

	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}
func (s *Store) Order() storage.OrderRepoI {

	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}

	return s.order
}

// GORM
// ROW
// SQLBUILDER
// SQLX
// PGXPOOL
