package postgres

import (
	"context"
	"errors"
	"market/internal/app/users"
	postgres "market/internal/repository/posgresql/db"
	"market/internal/utils"
	"market/internal/utils/security"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UsersRepo struct {
	db *postgres.Queries
}

func NewUsersRepo(db *postgres.Queries) users.Repository {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) SignIn(username, password string) (*users.User, error) {
	user, err := r.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	ok := security.CheckPassword(user.Password, password)
	if !ok {
		return nil, errors.New("invalid password")
	}
	return user, nil
}
func (r *UsersRepo) Create(u *users.NewUser) (*users.User, error) {
	hashPsw, err := security.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	username := utils.NormalizeUsername(u.Username)
	user, err := r.db.CreateUser(
		context.Background(),
		postgres.CreateUserParams{
			Username:     username,
			Name:         u.Name,
			Email:        &u.Email,
			PhoneNumber:  u.PhoneNumbers,
			PasswordHash: hashPsw,
			Data:         &u.Data,
			Role:         pareRoleToPG(u.Role),
		},
	)
	if err != nil {
		return nil, err
	}
	return &users.User{
		ID:           user.ID.String(),
		Username:     username,
		Name:         user.Name,
		Email:        user.Email,
		Data:         user.Data,
		PhoneNumbers: user.PhoneNumber,
		Role:         pareRolePgToRole(user.Role),
		CreatedAt:    user.Created.Time,
		UpdatedAt:    user.Updated.Time,
	}, nil
}

func (r *UsersRepo) GetByUsername(username string) (*users.User, error) {
	user, err := r.db.GetUserByUsername(
		context.Background(),
		username,
	)
	if err != nil {
		return nil, err
	}
	return &users.User{
		ID:           user.ID.String(),
		Username:     username,
		Name:         user.Name,
		Email:        user.Email,
		Data:         user.Data,
		Password:     user.PasswordHash,
		PhoneNumbers: user.PhoneNumber,
		Role:         pareRolePgToRole(user.Role),
		CreatedAt:    user.Created.Time,
		UpdatedAt:    user.Updated.Time,
	}, nil
}

func (r *UsersRepo) GetById(id string) (*users.User, error) {
	uuid, err := uuid.Parse(id)
	userId := pgtype.UUID{
		Bytes: uuid,
		Valid: true,
	}
	user, err := r.db.GetUserByid(
		context.Background(),
		userId,
	)
	if err != nil {
		return nil, err
	}
	return &users.User{
		ID:           user.ID.String(),
		Username:     user.Username,
		Name:         user.Name,
		Email:        user.Email,
		Data:         user.Data,
		Password:     user.PasswordHash,
		PhoneNumbers: user.PhoneNumber,
		Role:         pareRolePgToRole(user.Role),
		CreatedAt:    user.Created.Time,
		UpdatedAt:    user.Updated.Time,
	}, nil
}

func pareRoleToPG(role users.Role) postgres.NullUserRole {
	pgRole := postgres.NullUserRole{}
	switch role {
	case users.RoleAdmin:
		pgRole.UserRole = postgres.UserRoleAdmin
		pgRole.Valid = true
		break
	case users.RoleCashier:
		pgRole.UserRole = postgres.UserRoleCashier
		pgRole.Valid = true
		break
	case users.RoleClient:
		pgRole.UserRole = postgres.UserRoleClient
		pgRole.Valid = true
		break
	case users.RoleDealer:
		pgRole.UserRole = postgres.UserRoleDealer
		pgRole.Valid = true
		break
	}
	return pgRole
}

func pareRolePgToRole(role postgres.NullUserRole) users.Role {
	pgRole := users.RoleClient
	switch role.UserRole {
	case postgres.UserRoleAdmin:
		return users.RoleAdmin
	case postgres.UserRoleCashier:
		return users.RoleCashier
	case postgres.UserRoleClient:
		return users.RoleClient
	case postgres.UserRoleDealer:
		return users.RoleDealer
	}
	return pgRole
}
