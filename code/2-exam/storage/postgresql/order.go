package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) InfoOfSoldProductsByStaffer(ctx context.Context, req *models.GetListEmployeeRequest) (resp *models.GetListEmployeeResponse, err error) {

	resp = &models.GetListEmployeeResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
	SELECT 
	staf.first_name||' '||staf.last_name as employee,
	cate.category_name as category,
	p.product_name as product,
	(oi.quantity) AS total_amount,
	(oi.list_price) * (oi.quantity)  AS total_price,
	CAST(o.order_date::timestamp AS VARCHAR) as date 
	from orders as o
	join staffs as staf on staf.staff_id = o.staff_id 
	join stores as s on s.store_id = o.store_id
	join stocks as stoc on stoc.store_id = s.store_id
	join products as p on p.product_id = stoc.product_id
	join categories as cate on cate.category_id = p.category_id 
	join order_items as oi on oi.order_id = o.order_id
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
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	sum := 0
	for rows.Next() {

		var employee models.Employee

		err = rows.Scan(
			&employee.EmployeeName,
			// &employee.StoreName,
			&employee.CategoryName,
			&employee.ProductName,
			&employee.Total_amount,
			&employee.TotalPrice,
			&employee.Date,
		)
		if err != nil {
			return nil, err
		}

		resp.Employees = append(resp.Employees, &employee)
		sum++
	}
	resp.Count = sum

	return resp, nil
}

func (r *orderRepo) TotalPriceWithOrder(ctx context.Context, req *models.OrderTotalPrice) (string, error) {

	var (
		query       string
		PromoCode   models.PromoCode
		total_price float64
	)
	query = `
	SELECT 
	sum(oi.list_price) as total_price
	from orders as o 
	join order_items as oi on oi.order_id = o.order_id
	join products as p on p.product_id = oi.product_id
	where o.order_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.OrderId).Scan(&total_price)

	if total_price <= 0 {
		return "", fmt.Errorf("not available with this order")
	}

	if err != nil {
		return "", err
	}

	query = `
		SELECT
		name,
		discount,
		discount_type,
			order_limit_price
			FROM promo_code
			WHERE name = $1
			`

	var params = map[string]interface{}{
		"name":               req.Name,
		"discount":           req.Discount,
		"discount_type":      req.DiscountType,
		"order_limit_price ": req.OrderLimitPrice,
	}
	err = r.db.QueryRow(ctx, query, params).Scan(
		&PromoCode.Name,
		&PromoCode.Discount,
		&PromoCode.DiscountType,
		&PromoCode.OrderLimitPrice,
	)
	fmt.Println("ccccccccccccccccccccccc", req.Name)
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", err)
	if err != nil {
		return "", err
	}

	if total_price < 0 {
		return "", fmt.Errorf("don't have money")
	}

	if PromoCode.DiscountType == "fix" && PromoCode.OrderLimitPrice < total_price {
		total_price = total_price - PromoCode.Discount
	} else if PromoCode.DiscountType == "percent" && PromoCode.OrderLimitPrice < total_price {
		total_price = total_price - total_price*PromoCode.Discount/100
	}
	return fmt.Sprintf("total_price: %v", total_price), nil
}

func (r *orderRepo) Create(ctx context.Context, req *models.CreateOrder) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO orders(
			order_id, 
			customer_id, 
			order_status,
			order_date,
			required_date,
			shipped_date,
			store_id,
			staff_id
		)
		VALUES (
			(
				SELECT MAX(order_id) + 1 FROM orders
			)
			, $1, $2, now()::date, $3, $4, $5, $6) RETURNING order_id
	`
	fmt.Println(query)

	err := r.db.QueryRow(ctx, query,
		helper.NewNullInt32(req.CustomerId),
		req.OrderStatus,
		req.RequiredDate,
		helper.NewNullString(req.ShippedDate),
		req.StoreId,
		req.StaffId,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *orderRepo) GetByID(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error) {

	var (
		query string
		order models.Order
	)

	query = `
		WITH order_item_data AS (
			SELECT
				oi.order_id AS order_id,
				JSONB_AGG (
					JSONB_BUILD_OBJECT (
						'order_id', oi.order_id,
						'item_id', oi.item_id,
						'product_id', oi.product_id,
						'quantity', oi.quantity,
						'list_price', oi.list_price,
						'discount', oi.discount,
						'cell_price', oi.cell_price
					)
				) AS order_items
		
			FROM order_items AS oi
			WHERE oi.order_id = $1
			GROUP BY oi.order_id
		)
		SELECT
			o.order_id, 
			o.customer_id,
		
			c.customer_id,
			c.first_name,
			c.last_name,
			COALESCE(c.phone, ''),
			c.email,
			COALESCE(c.street, ''),
			COALESCE(c.city, ''),
			COALESCE(c.state, ''),
			COALESCE(c.zip_code, 0),
			
			o.order_status,
			CAST(o.order_date::timestamp AS VARCHAR),
			CAST(o.required_date::timestamp AS VARCHAR),
			COALESCE(CAST(o.shipped_date::timestamp AS VARCHAR), ''),
			o.store_id,
		
			s.store_id,
			s.store_name,
			COALESCE(s.phone, ''),
			COALESCE(s.email, ''),
			COALESCE(s.street, ''),
			COALESCE(s.city, ''),
			COALESCE(s.state, ''),
			COALESCE(s.zip_code, ''),
		
			o.staff_id,
			st.staff_id,
			st.first_name,
			st.last_name,
			st.email,
			COALESCE(st.phone, ''),
			st.active,
			st.store_id,
			COALESCE(st.manager_id, 0),
		
			oi.order_items
		
		FROM orders AS o
		JOIN customers AS c ON c.customer_id = o.customer_id
		JOIN stores AS s ON s.store_id = o.store_id
		JOIN staffs AS st ON st.staff_id = o.staff_id
		JOIN order_item_data AS oi ON oi.order_id = o.order_id
		WHERE o.order_id = $1	
	`
	order.CustomerData = &models.Customer{}
	order.StoreData = &models.Store{}
	order.StaffData = &models.Staff{}
	orderItemObject := pgtype.JSON{}

	err := r.db.QueryRow(ctx, query, req.OrderId).Scan(
		&order.OrderId,
		&order.CustomerId,
		&order.CustomerData.CustomerId,
		&order.CustomerData.FirstName,
		&order.CustomerData.LastName,
		&order.CustomerData.Phone,
		&order.CustomerData.Email,
		&order.CustomerData.Street,
		&order.CustomerData.City,
		&order.CustomerData.State,
		&order.CustomerData.ZipCode,

		&order.OrderStatus,
		&order.OrderDate,
		&order.RequiredDate,
		&order.ShippedDate,

		&order.StoreId,

		&order.StoreData.StoreId,
		&order.StoreData.StoreName,
		&order.StoreData.Phone,
		&order.StoreData.Email,
		&order.StoreData.Street,
		&order.StoreData.City,
		&order.StoreData.State,
		&order.StoreData.ZipCode,
		&order.StaffId,
		&order.StaffData.StaffId,
		&order.StaffData.FirstName,
		&order.StaffData.LastName,
		&order.StaffData.Email,
		&order.StaffData.Phone,
		&order.StaffData.Active,
		&order.StaffData.StoreId,
		&order.StaffData.ManagerId,

		&orderItemObject,
	)
	if err != nil {
		return nil, err
	}

	orderItemObject.AssignTo(&order.OrderItems)

	return &order, nil
}

func (r *orderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error) {

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
			o.order_id, 
			o.customer_id,
			c.customer_id,

			c.first_name,
			c.last_name,
			COALESCE(c.phone, ''),
			c.email,
			COALESCE(c.street, ''),
			COALESCE(c.city, ''),
			COALESCE(c.state, ''),
			COALESCE(c.zip_code, 0),

			o.order_status,
			CAST(o.order_date::timestamp AS VARCHAR),
			CAST(o.required_date::timestamp AS VARCHAR),
			COALESCE(CAST(o.shipped_date::timestamp AS VARCHAR), ''),
			o.store_id,

			s.store_id,
			s.store_name,
			COALESCE(s.phone, ''),
			COALESCE(s.email, ''),
			COALESCE(s.street, ''),
			COALESCE(s.city, ''),
			COALESCE(s.state, ''),
			COALESCE(s.zip_code, ''),

			o.staff_id,
			st.staff_id,
			st.first_name,
			st.last_name,
			st.email,
			COALESCE(st.phone, ''),
			st.active,
			st.store_id,
			COALESCE(st.manager_id, 0)

		FROM orders AS o
		JOIN customers AS c ON c.customer_id = o.customer_id
		JOIN stores AS s ON s.store_id = o.store_id
		JOIN staffs AS st ON st.staff_id = o.staff_id
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
		var order models.Order
		order.CustomerData = &models.Customer{}
		order.StoreData = &models.Store{}
		order.StaffData = &models.Staff{}

		err = rows.Scan(
			&resp.Count,
			&order.OrderId,
			&order.CustomerId,
			&order.CustomerData.CustomerId,
			&order.CustomerData.FirstName,
			&order.CustomerData.LastName,
			&order.CustomerData.Phone,
			&order.CustomerData.Email,
			&order.CustomerData.Street,
			&order.CustomerData.City,
			&order.CustomerData.State,
			&order.CustomerData.ZipCode,

			&order.OrderStatus,
			&order.OrderDate,
			&order.RequiredDate,
			&order.ShippedDate,

			&order.StoreId,

			&order.StoreData.StoreId,
			&order.StoreData.StoreName,
			&order.StoreData.Phone,
			&order.StoreData.Email,
			&order.StoreData.Street,
			&order.StoreData.City,
			&order.StoreData.State,
			&order.StoreData.ZipCode,
			&order.StaffId,
			&order.StaffData.StaffId,
			&order.StaffData.FirstName,
			&order.StaffData.LastName,
			&order.StaffData.Email,
			&order.StaffData.Phone,
			&order.StaffData.Active,
			&order.StaffData.StoreId,
			&order.StaffData.ManagerId,
		)
		if err != nil {
			return nil, err
		}

		resp.Orders = append(resp.Orders, &order)
	}

	return resp, nil
}

func (r *orderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		orders
		SET
			order_id = :order_id, 
			customer_id = :customer_id, 
			order_status = :order_status, 
			order_date = :order_date,
			required_date = :required_date,
			shipped_date = :shipped_date,
			store_id = :store_id,
			staff_id = :staff_id
		WHERE order_id = :order_id
	`

	params = map[string]interface{}{
		"order_id":      req.OrderId,
		"customer_id":   req.CustomerId,
		"order_status":  req.OrderStatus,
		"order_date":    req.OrderDate,
		"required_date": req.RequiredDate,
		"shipped_date":  helper.NewNullString(req.ShippedDate),
		"store_id":      req.StoreId,
		"staff_id":      req.StaffId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *orderRepo) UpdatePatch(ctx context.Context, req *models.PatchRequest) (int64, error) {
	var (
		query string
		set   string
	)

	if len(req.Fields) <= 0 {
		return 0, errors.New("no fields")
	}

	i := 0
	for key := range req.Fields {
		if i == len(req.Fields)-1 {
			set += fmt.Sprintf(" %s = :%s ", key, key)
		} else {
			set += fmt.Sprintf(" %s = :%s, ", key, key)
		}
		i++
	}

	query = `
		UPDATE
		orders
		SET
		` + set + `
		WHERE order_id = :order_id
	`

	req.Fields["order_id"] = req.ID

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	fmt.Println(query)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *orderRepo) Delete(ctx context.Context, req *models.OrderPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM orders
		WHERE order_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.OrderId)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

// ------------------------------------------------------------------------------------------------------------
func (r *orderRepo) AddOrderItem(ctx context.Context, req *models.CreateOrderItem) error {

	req.Cell_price = req.ListPrice - float64(req.ListPrice)*req.Discount

	query := `
		INSERT INTO order_items(
			order_id, 
			item_id, 
			product_id,
			quantity,
			list_price,
			discount,
			cell_price
		)
		VALUES (
			$1, 
			(
				SELECT COALESCE(MAX(item_id), 0) + 1 FROM order_items WHERE order_id = $1
			)
			, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(ctx, query,
		req.OrderId,
		req.ProductId,
		req.Quantity,
		req.ListPrice,
		req.Discount,
		req.Cell_price,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepo) RemoveOrderItem(ctx context.Context, req *models.OrderItemPrimaryKey) error {

	query := `
		DELETE FROM order_items WHERE order_id = $1 AND item_id = $2
	`
	_, err := r.db.Exec(ctx, query, req.OrderId, req.ItemId)

	if err != nil {
		return err
	}

	return nil
}
