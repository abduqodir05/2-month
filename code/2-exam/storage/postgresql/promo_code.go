package postgresql

import (
	"app/api/models"
	// "app/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PromoCodeRepo struct {
	db *pgxpool.Pool
}

func NewPromoCodeRepo(db *pgxpool.Pool) *PromoCodeRepo {
	return &PromoCodeRepo{
		db: db,
	}
}

func (r *PromoCodeRepo) CreatePromoCode(ctx context.Context, req *models.CreatePromoCode) (string, error) {
	var (
		query string
		promoName    string
	)

	query = `
		INSERT INTO promo_code(
			name,
			discount,
			discount_type,
			order_limit_price
		)
			Values (
				$1, $2, $3, $4) RETURNING name
	`

	err := r.db.QueryRow(ctx, query,
		req.Name,
		req.Discount,
		req.DiscountType,
		req.OrderLimitPrice,
	).Scan(&promoName)

	if err != nil {
		return "", err
	}

	return promoName, nil
}

func (r *PromoCodeRepo) GetByIDPromoCode(ctx context.Context, req *models.PromoCodePrimaryKey) (*models.PromoCode, error) {

	var (
		query     string
		PromoCode models.PromoCode
	)

	query = `
		SELECT
		name,
		discount,
		discount_type,
		order_limit_price
		FROM promo_code 
		WHERE name = $1
	`

	err := r.db.QueryRow(ctx, query, req.Name).Scan(
		&PromoCode.Name,
		&PromoCode.Discount,
		&PromoCode.DiscountType,
		&PromoCode.OrderLimitPrice,
	)
	if err != nil {
		return nil, err
	}

	return &PromoCode, nil
}

func (r *PromoCodeRepo) GetListPromoCode(ctx context.Context, req *models.GetListPromoCodeRequest) (resp *models.GetListPromoCodeResponse, err error) {

	resp = &models.GetListPromoCodeResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			name,
			discount,
			discount_type,
			order_limit_price
		FROM promo_code 
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
		var PromoCode models.PromoCode
		err = rows.Scan(
			&resp.Count,
			&PromoCode.Name,
			&PromoCode.Discount,
			&PromoCode.DiscountType,
			&PromoCode.OrderLimitPrice,
		)
		if err != nil {
			return nil, err
		}

		resp.PromoCodes = append(resp.PromoCodes, &PromoCode)
	}

	return resp, nil
}


func (r *PromoCodeRepo) DeletePromoCode(ctx context.Context, req *models.PromoCodePrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM promo_code
		WHERE name = $1
	`

	result, err := r.db.Exec(ctx, query, req.Name)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
