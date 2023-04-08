package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}
func (r UserRepo) CreateUser(ctx context.Context, req *models.CreateUser) (string, error) {
	var (
		query string
		id    = uuid.New().String()
	)

	query = `
		INSERT INTO users(
			id,
			first_name, 
			last_name,
			login,
			password,
			phone_number,
			updated_at,
			created_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, now(), now())
		`
	_, err := r.db.Exec(ctx, query,
		id,
		helper.NewNullString(req.FirstName),
		helper.NewNullString(req.LastName),
		helper.NewNullString(req.Login),
		helper.NewNullString(req.Password),
		helper.NewNullString(req.PhoneNumber),
	)
	if err != nil {
		return "", err
	}
	fmt.Println("Expexting UUID>>>>>", id)

	return id, nil
}

func (r *UserRepo) GetByIDUser(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query string
		User  models.User
	)

	query = `
		SELECT
			id,
			first_name, 
			last_name,
			login,
			password,
			phone_number
		FROM users
		WHERE id = $1
		`

	if len(req.Login) > 0 {
		err := r.db.QueryRow(ctx, `select id from users where login = $1`, req.Login).Scan(&req.UserId)

		if err != nil {
			return nil, err
		}
	}

	err := r.db.QueryRow(ctx, query, req.UserId).Scan(
		&User.UserId,
		&User.FirstName,
		&User.LastName,
		&User.Login,
		&User.Password,
		&User.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return &User, nil
}

func (r *UserRepo) GetListUser(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) {

	resp = &models.GetListUserResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			first_name, 
			last_name,
			login,
			password,
			phone_number
		FROM users
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var User models.User
		err = rows.Scan(
			&resp.Count,
			&User.UserId,
			&User.FirstName,
			&User.LastName,
			&User.Login,
			&User.Password,
			&User.PhoneNumber,
		)
		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users)
	}

	return resp, nil
}
