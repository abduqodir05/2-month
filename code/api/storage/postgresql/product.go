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

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) CreateProduct(ctx context.Context, req *models.CreateProduct) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO products(
			id, 
			name, 
			price,
			category_id, 
			updated_at
		)
		VALUES ($1, $2, $3, now())
	`

	_, err := r.db.Exec(ctx, query,
		id.String(),
		req.Name,
		req.Price,
		req.CategoryId,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *ProductRepo) GetByIdProduct(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {

	var (
		query string
		Product  models.Product
	)

	query = `
		SELECT
			id,
			name,
			price,
			category_id,
			created_at,
			updated_at
		FROM products
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&Product.Id,
		&Product.Name,
		&Product.Price,
		&Product.CategoryId,
		&Product.CreatedAt,
		&Product.UpdatedAt,
	)
	fmt.Println(Product)
	if err != nil {
		return nil, err
	}

	return &Product, nil
}

func (r *ProductRepo) GetListProduct(ctx context.Context, req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error) {

	resp = &models.GetListProductResponse{}

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
			price,
			category_id,
			created_at,
			updated_at
		FROM products
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

		var products models.Product
		err = rows.Scan(
			&resp.Count,
			&products.Id,
			&products.Name,
			&products.Price,
			&products.CategoryId,
			&products.CreatedAt,
			&products.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &products)
	}

	return resp, nil
}

func (r *ProductRepo) UpdateProduct(ctx context.Context, req *models.UpdateProduct) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			products
		SET
			name = :name,
			price = :price,
			category_id,= :category_id,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"name":         req.Name,
		"price":		req.Price,
		"category_id":  req.CategoryId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *ProductRepo) PatchProduct(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			products
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

func (r *ProductRepo) DeleteProduct(ctx context.Context, req *models.ProductPrimaryKey) error {

	_, err := r.db.Exec(ctx,
		"DELETE FROM products WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
