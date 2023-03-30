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

type CustomerRepo struct {
	db *pgxpool.Pool
}

func NewCustomerRepo(db *pgxpool.Pool) *CustomerRepo {
	return &CustomerRepo{
		db: db,
	}
}

func (r *CustomerRepo) CreateCustomer(ctx context.Context, req *models.CreateCustomer) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO Customers(
			id, 
			name, 
			phone, 
			updated_at
		)
		VALUES ($1, $2, $3, now())
	`

	_, err := r.db.Exec(ctx, query,
		id.String(),
		req.Name,
		req.Phone,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *CustomerRepo) GetByIdCustomer(ctx context.Context, req *models.CustomerPrimaryKey) (*models.Customer, error) {

	var (
		query string
		Customer  models.Customer
	)

	query = `
		SELECT
			id,
			name,
			phone,
			created_at,
			updated_at
		FROM customers
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&Customer.Id,
		&Customer.Name,
		&Customer.Phone,
		&Customer.CreatedAt,
		&Customer.UpdatedAt,
	)
	fmt.Println(Customer)
	if err != nil {
		return nil, err
	}

	return &Customer, nil
}

func (r *CustomerRepo) GetListCustomer(ctx context.Context, req *models.GetListCustomerRequest) (resp *models.GetListCustomerResponse, err error) {

	resp = &models.GetListCustomerResponse{}

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
			phone,
			created_at,
			updated_at
		FROM customers
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

		var Customers models.Customer
		err = rows.Scan(
			&resp.Count,
			&Customers.Id,
			&Customers.Name,
			&Customers.Phone,
			&Customers.CreatedAt,
			&Customers.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Customers = append(resp.Customers, &Customers)
	}

	return resp, nil
}

func (r *CustomerRepo) UpdateCustomer(ctx context.Context, req *models.UpdateCustomer) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			Customers
		SET
			name = :name,
			phone = :phone,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"name":         req.Name,
		"phone": req.Phone,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *CustomerRepo) PatchCustomer(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			Customers
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

func (r *CustomerRepo) DeleteCustomer(ctx context.Context, req *models.CustomerPrimaryKey) error {

	_, err := r.db.Exec(ctx,
		"DELETE FROM customers WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
