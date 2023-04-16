package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"gitlab.com/book6485215/user_go_user_service/genproto/user_service"
	"gitlab.com/book6485215/user_go_user_service/models"
	"gitlab.com/book6485215/user_go_user_service/pkg/helper"
	"gitlab.com/book6485215/user_go_user_service/storage"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.UserRepoI {
	return &UserRepo{
		db: db,
	}
}

func (c *UserRepo) Create(ctx context.Context, req *user_service.CreateUser) (resp *user_service.UserPrimaryKey, err error) {

	var id = uuid.New()

	query := `INSERT INTO "users" (
				id,
				first_name,
				last_name,
				type,
				updated_at
			) VALUES ($1, $2, $3, $4, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.FirstName,
		req.LastName,
		req.Type,
	)

	if err != nil {
		return nil, err
	}

	return &user_service.UserPrimaryKey{Id: id.String()}, nil
}

func (c *UserRepo) GetByPKey(ctx context.Context, req *user_service.UserPrimaryKey) (resp *user_service.User, err error) {

	query := `
		SELECT
			id,
			first_name,
			last_name,
			type,
			created_at,
			updated_at
		FROM "users"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		firstName sql.NullString
		lastName  sql.NullString
		userType  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&firstName,
		&lastName,
		&userType,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &user_service.User{
		Id:        id.String,
		FirstName: firstName.String,
		LastName:  lastName.String,
		Type:      userType.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *UserRepo) GetAll(ctx context.Context, req *user_service.GetListUserRequest) (resp *user_service.GetListUserResponse, err error) {

	resp = &user_service.GetListUserResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE"
		sort   = " ORDER BY created_at DESC"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			first_name,
			last_name,
			type,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "users"
	`

	if req.GetLimit() > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}

	if req.GetOffset() > 0 {
		offset = " OFFSET :offset"
		params["offset"] = req.Offset
	}

	query += filter + sort + offset + limit

	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := c.db.Query(ctx, query, args...)
	defer rows.Close()

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			firstName sql.NullString
			lastName  sql.NullString
			userType  sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&firstName,
			&lastName,
			&userType,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Users = append(resp.Users, &user_service.User{
			Id:        id.String,
			FirstName: firstName.String,
			LastName:  lastName.String,
			Type:      userType.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *UserRepo) Update(ctx context.Context, req *user_service.UpdateUser) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "users"
			SET
				first_name = :first_name,
				last_name = :last_name,
				type = :type,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":         req.GetId(),
		"first_name": req.GetFirstName(),
		"last_name":  req.GetLastName(),
		"type":       req.GetType(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *UserRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

	var (
		set   = " SET "
		ind   = 0
		query string
	)

	if len(req.Fields) == 0 {
		err = errors.New("no updates provided")
		return
	}

	req.Fields["id"] = req.Id

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s ", key, key)
		if ind != len(req.Fields)-1 {
			set += ", "
		}
		ind++
	}

	query = `
		UPDATE
			"users"
	` + set + ` , updated_at = now()
		WHERE
			id = :id
	`

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), err
}

func (c *UserRepo) Delete(ctx context.Context, req *user_service.UserPrimaryKey) error {

	query := `DELETE FROM "users" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
