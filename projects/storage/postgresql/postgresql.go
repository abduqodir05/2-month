package postgresql

import (
	"app/config"
	"app/storage"
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
)

type Store struct {
	db   *sql.DB
	user storage.UserRepoI
}

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {

	connection := fmt.Sprintf(
		"host=%s user=%s database=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)
	db, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Store{
		db:   db,
		user: NewUserRepo(db),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}
	return s.user
}
