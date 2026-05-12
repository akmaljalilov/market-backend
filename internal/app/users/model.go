package users

import (
	"time"
)

type Role int

const (
	RoleAdmin Role = iota + 1
	RoleCashier
	RoleClient
	RoleDealer
)

type SignInResponse struct {
	User
	Token string `json:"token"`
}
type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Name         string    `json:"name"`
	Email        *string   `json:"email"`
	Data         *string   `json:"data"`
	PhoneNumbers []string  `json:"phoneNumbers"`
	Password     string    `json:"-"`
	Role         Role      `json:"role"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type NewUserResp struct {
	UUID     string
	Username string
}
type NewUser struct {
	Username     string
	Name         string
	Email        string
	Password     string
	Data         string
	PhoneNumbers []string
	Role         Role
}
