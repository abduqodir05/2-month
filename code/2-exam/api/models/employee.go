package models

type Employee struct {
	EmployeeName string  `json:"employee"`
	CategoryName string  `json:"category"`
	ProductName  string  `json:"product"`
	Total_amount     int     `json:"total_amount`
	TotalPrice   float64 `json:"total_price"`
	Date         string  `json:"date"`
}

type GetListEmployeeRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListEmployeeResponse struct {
	Count     int         `json:"count"`
	Employees []*Employee `json:"employee"`
}
