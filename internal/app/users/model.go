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

type UserID string

type User struct {
	ID        UserID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewUser struct {
	Name         string
	Email        string
	Password     string
	Data         string
	PhoneNumbers []string
	Role         Role
}
