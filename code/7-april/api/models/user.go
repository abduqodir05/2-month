package models

type User struct {
	UserId      string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Register    string `json:"register"`
	Login       string `json:"login"`
	Password    string `json:"password"`
}

type UserPrimaryKey struct {
	UserId string `json:"id"`
	Login  string `json:"login"`
}

type CreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Login       string `json:"login"`
}

type GetListUserRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListUserResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"Users"`
}
