package auth

import "market/internal/app/users"

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	users.SignInResponse
}
