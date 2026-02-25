package postgres

import (
	"context"
	"market/internal/app/users"
	postgres "market/internal/repository/posgresql/db"
	"market/internal/utils/security"
)

type UsersRepo struct {
	db *postgres.Queries
}

func NewUsersRepo(db *postgres.Queries) users.Repository {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) Create(u *users.NewUser) (string, error) {
	hashPsw, err := security.HashPassword(u.Password)
	if err != nil {
		return "", err
	}
	id, err := r.db.CreateUser(
		context.Background(),
		postgres.CreateUserParams{
			Name:         u.Name,
			Email:        &u.Email,
			PhoneNumber:  u.PhoneNumbers,
			PasswordHash: hashPsw,
			Data:         &u.Data,
			Role:         pareRoleToPG(u.Role),
		},
	)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func pareRoleToPG(role users.Role) postgres.NullUserRole {
	pgRole := postgres.NullUserRole{}
	switch role {
	case users.RoleAdmin:
		pgRole.UserRole = postgres.UserRoleAdmin
		break
	case users.RoleCashier:
		pgRole.UserRole = postgres.UserRoleCashier
		break
	case users.RoleClient:
		pgRole.UserRole = postgres.UserRoleClient
		break
	case users.RoleDealer:
		pgRole.UserRole = postgres.UserRoleDealer
		break
	}
	return pgRole
}
