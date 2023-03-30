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

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) CreateCategory(ctx context.Context, req *models.CreateCategory) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO categories(
			id, 
			name, 
			parent_id, 
			updated_at
		)
		VALUES ($1, $2, $3, now())
	`

	_, err := r.db.Exec(ctx, query,
		id.String(),
		req.Name,
		req.ParentId,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *CategoryRepo) GetByIdCategory(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error) {

	var (
		query string
		Category  models.Category
	)

	query = `
		SELECT
			id,
			name,
			parent_id,
			created_at,
			updated_at
		FROM categories
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&Category.Id,
		&Category.Name,
		&Category.ParentId,
		&Category.CreatedAt,
		&Category.UpdatedAt,
	)
	fmt.Println(Category)
	if err != nil {
		return nil, err
	}

	return &Category, nil
}

func (r *CategoryRepo) GetListCategory(ctx context.Context, req *models.GetListCategoryRequest) (resp *models.GetListCategoryResponse, err error) {

	resp = &models.GetListCategoryResponse{}

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
			parent_id,
			created_at,
			updated_at
		FROM categories
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

		var categories models.Category
		err = rows.Scan(
			&resp.Count,
			&categories.Id,
			&categories.Name,
			&categories.ParentId,
			&categories.CreatedAt,
			&categories.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Categories = append(resp.Categories, &categories)
	}

	return resp, nil
}

func (r *CategoryRepo) UpdateCategory(ctx context.Context, req *models.UpdateCategory) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			categories
		SET
			name = :name,
			parent_id = :parent_id,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"name":         req.Name,
		"parent_id": req.ParentId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *CategoryRepo) PatchCategory(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			categories
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

func (r *CategoryRepo) DeleteCategory(ctx context.Context, req *models.CategoryPrimaryKey) error {

	_, err := r.db.Exec(ctx,
		"DELETE FROM categories WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
