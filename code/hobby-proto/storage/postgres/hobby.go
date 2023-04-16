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

type HobbyRepo struct {
	db *pgxpool.Pool
}

func NewHobbyRepo(db *pgxpool.Pool) storage.HobbyRepoI {
	return &HobbyRepo{
		db: db,
	}
}

func (c *HobbyRepo) CreateHobby(ctx context.Context, req *user_service.CreateHobby) (resp *user_service.HobbyPrimaryKey, err error) {

	var id = uuid.New()
	// log.Info("that's id from storage", logger.String("id", id))
fmt.Println("----------------------------")
	query := `INSERT INTO "hobbies" (
				id,
				name,
				updated_at
			) VALUES ($1, $2, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.Name,
	)

	if err != nil {
		fmt.Println("create hobby error:>>>>", err)
		return nil, err
	}

	return &user_service.HobbyPrimaryKey{Id: id.String()}, nil
}

func (c *HobbyRepo) GetByPKeyHobby(ctx context.Context, req *user_service.HobbyPrimaryKey) (resp *user_service.Hobby, err error) {

	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM "hobbies"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		name      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &user_service.Hobby{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *HobbyRepo) GetAllHobby(ctx context.Context, req *user_service.GetListHobbyRequest) (resp *user_service.GetListHobbyResponse, err error) {

	resp = &user_service.GetListHobbyResponse{}

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
			name,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "hobbies"
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

	if err != nil {
		fmt.Println("here get list hobby error:>>>", err)
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        sql.NullString
			name      sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Hobbies = append(resp.Hobbies, &user_service.Hobby{
			Id:        id.String,
			Name:      name.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *HobbyRepo) UpdateHobby(ctx context.Context, req *user_service.UpdateHobby) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "hobbies"
			SET
				name = : name,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"name": req.GetName(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *HobbyRepo) UpdatePatchHobby(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

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
			"hobbies"
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

func (c *HobbyRepo) DeleteHobby(ctx context.Context, req *user_service.HobbyPrimaryKey) error {

	query := `DELETE FROM "hobbies" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
