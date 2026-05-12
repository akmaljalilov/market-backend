package user

import "market/internal/app/users"

type createUserResponse struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}
type createUserReq struct {
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	PhoneNumbers []string   `json:"phoneNumbers"`
	Role         users.Role `json:"role"`
	Data         string     `json:"data"`
}
type currentUserResponse struct {
	users.User
}
