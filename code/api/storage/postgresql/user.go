package postgresql

import (
	"context"
	// "database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, req *models.CreateUser) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO users(
			id, 
			name, 
			phone_number, 
			updated_at
		)
		VALUES ($1, $2, $3, now())
	`

	_, err := r.db.Exec(ctx, query,
		id.String(),
		req.Name,
		req.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *userRepo) GetByIdUser(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query string
		user  models.User
	)

	query = `
		SELECT
			id,
			name,
			phone_number,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&user.Id,
		&user.Name,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetListUser(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) {

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
			name,
			phone_number,
			created_at,
			updated_at
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

		var users models.User
		err = rows.Scan(
			&resp.Count,
			&users.Id,
			&users.Name,
			&users.PhoneNumber,
			&users.CreatedAt,
			&users.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &users)
	}

	return resp, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, req *models.UpdateUser) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			users
		SET
			name = :name,
			phone_number = :phone_number,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"name":         req.Name,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *userRepo) PatchUser(ctx context.Context, req *models.PatchRequest) (int64, error) {

	var (
		query string
		set   string
	)

	if len(req.Fields) <= 0 {
		return 0, errors.New("no fields")
	}

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s, ", key, key)
	}

	query = `
		UPDATE
			users
		SET
	` + set + ` updated_at = now()
		WHERE id = :id
	`

	req.Fields["id"] = req.ID

	fmt.Println(req.Fields)

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *userRepo) DeleteUser(ctx context.Context, req *models.UserPrimaryKey) error {

	_, err := r.db.Exec(ctx,
		"DELETE FROM users WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
