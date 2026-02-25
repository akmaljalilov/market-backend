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

type User struct {
	ID           string
	Username     string
	Name         string
	Email        *string
	Data         *string
	PhoneNumbers []string
	Role         Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
