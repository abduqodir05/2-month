package postgresql

import (
	"app/api/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(req *models.CreateUser) (string, error) {
	
	var (
		query string
		id    = uuid.New()
	)
	query = `
	INSERT INTO users(
		id,
		name,
		phone
	)
	VALUES ($1, $2, $3)
		`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
		req.Phone,
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *UserRepo) UpdateUser(req *models.UpdateUser) (string, error) {

	query := `
	Update users set name = $1, phone = $2
	where id = $3
	`

	_, err := r.db.Exec(query,
		req.Name,
		req.Phone,
		req.Id,
	)

	if err != nil {
		return "nil", err
	}

	return req.Id, nil
}

func (r *UserRepo) DeleteUser(req *models.DeleteUser) error {

	query := `
	delete from users 
	where id = $1
	`

	_, err := r.db.Exec(query,
		req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetByIDUser(req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query string
		User  models.User
	)

	query = `
		SELECT
			id,
			name,
			phone
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&User.Id,
		&User.Name,
		&User.Phone,
	)

	if err != nil {
		return nil, err
	}

	return &User, nil
}

func (r *UserRepo) GetListUser(req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) {

	resp = &models.GetListUserResponse{}

	var (
		query  string
		// filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			phone
		FROM users
	`

	// if len(req.Search) > 0 {
	// 	filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	// }

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit

	rows, err := r.db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var User models.User
		err = rows.Scan(
			&resp.Count,
			&User.Id,
			&User.Name,
			&User.Phone,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &User)
	}

	return resp, nil
}
