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

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, req *models.CreateOrder) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO orders(
			id, 
			name, 
			price,
			phone_number,
			latitude,
			longtitude,
			user_id,
			customer_id,
			courier_id,
			product_id,
			quantity, 
			updated_at
		)
		VALUES ($1, $2, $3, now())
	`

	_, err := r.db.Exec(ctx, query,
		id.String(),
		req.Name,
		req.Price,
		req.PhoneNumber,
		req.Latitude,
		req.Longtitude,
		req.UserId,
		req.CustomerId,
		req.CourierId,
		req.ProductId,
		req.Quantity,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *OrderRepo) GetByIdOrder(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error) {

	var (
		query string
		Order models.Order
	)

	query = `
		SELECT
		id, 
		name, 
		price,
		phone_number,
		latitude,
		longtitude,
		user_id,
		customer_id,
		courier_id,
		product_id,
		quantity, 
			created_at,
			updated_at
		FROM orders
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&Order.Id,
		&Order.Name,
		&Order.Price,
		&Order.PhoneNumber,
		&Order.Latitude,
		&Order.Longtitude,
		&Order.UserId,
		&Order.CustomerId,
		&Order.CourierId,
		&Order.ProductId,
		&Order.Quantity,
		&Order.CreatedAt,
		&Order.UpdatedAt,
	)
	fmt.Println(Order)
	if err != nil {
		return nil, err
	}

	return &Order, nil
}

func (r *OrderRepo) GetListOrder(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error) {

	resp = &models.GetListOrderResponse{}

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
		phone_number,
		latitude,
		longtitude,
		user_id,
		customer_id,
		courier_id,
		product_id,
		quantity, 
			created_at,
			updated_at
		FROM orders
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

		var orders models.Order
		err = rows.Scan(
			&resp.Count,
			&orders.Id,
			&orders.Name,
			&orders.Price,
			&orders.PhoneNumber,
			&orders.Latitude,
			&orders.Longtitude,
			&orders.UserId,
			&orders.CustomerId,
			&orders.CourierId,
			&orders.ProductId,
			&orders.Quantity,
			&orders.CreatedAt,
			&orders.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Orders = append(resp.Orders, &orders)
	}

	return resp, nil
}

func (r *OrderRepo) UpdateOrder(ctx context.Context, req *models.UpdateOrder) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			orders
		SET
			name = :name,
			price = :price,
			latitude= :latitude,
			longtitude= :longtitude,
			user_id= :user_id,
			customer_id= :customer_id,
			courier_id= :courier_id,
			product_id= :product_id,
			quantity= :quantity
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"name":        req.Name,
		"price":       req.Price,
		"latitude":    req.Latitude,
		"longtitude":  req.Longtitude,
		"user_id":     req.UserId,
		"customer_id": req.CustomerId,
		"courier_id":  req.CourierId,
		"product_id":  req.ProductId,
		"quantity":    req.Quantity,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *OrderRepo) PatchOrder(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			orders
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

func (r *OrderRepo) DeleteOrder(ctx context.Context, req *models.OrderPrimaryKey) error {

	_, err := r.db.Exec(ctx,
		"DELETE FROM orders WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
