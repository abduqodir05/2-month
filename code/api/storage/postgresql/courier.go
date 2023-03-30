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

type CourierRepo struct {
	db *pgxpool.Pool
}

func NewCourierRepo(db *pgxpool.Pool) *CourierRepo {
	return &CourierRepo{
		db: db,
	}
}

func (r *CourierRepo) CreateCourier(ctx context.Context, req *models.CreateCourier) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO couriers(
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

func (r *CourierRepo) GetByIdCourier(ctx context.Context, req *models.CourierPrimaryKey) (*models.Courier, error) {

	var (
		query string
		Courier  models.Courier
	)

	query = `
		SELECT
			id,
			name,
			phone_number,
			created_at,
			updated_at
		FROM couriers
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&Courier.Id,
		&Courier.Name,
		&Courier.PhoneNumber,
		&Courier.CreatedAt,
		&Courier.UpdatedAt,
	)
	fmt.Println(Courier)
	if err != nil {
		return nil, err
	}

	return &Courier, nil
}

func (r *CourierRepo) GetListCourier(ctx context.Context, req *models.GetListCourierRequest) (resp *models.GetListCourierResponse, err error) {

	resp = &models.GetListCourierResponse{}

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
		FROM couriers
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

		var couriers models.Courier
		err = rows.Scan(
			&resp.Count,
			&couriers.Id,
			&couriers.Name,
			&couriers.PhoneNumber,
			&couriers.CreatedAt,
			&couriers.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Couriers = append(resp.Couriers, &couriers)
	}

	return resp, nil
}

func (r *CourierRepo) UpdateCourier(ctx context.Context, req *models.UpdateCourier) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			couriers
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

func (r *CourierRepo) PatchCourier(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
			couriers
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

func (r *CourierRepo) DeleteCourier(ctx context.Context, req *models.CourierPrimaryKey) error {

	_, err := r.db.Exec(ctx,
		"DELETE FROM couriers WHERE id = $1", req.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
