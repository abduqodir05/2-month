package models

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Register string `json:"register"`
}
