package users

import (
	"context"
	users "market/internal/app/users"
	"market/internal/config"
	repo "market/internal/repository/posgresql"
	postgres "market/internal/repository/posgresql/db"
	"market/internal/services"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var userService *users.App

func init() {
	cfg := config.Postgres{
		Username: "root",
		Host:     "localhost",
		Password: "mysecretpassword",
		DbName:   "aj_market",
	}
	testDB, err := pgxpool.New(context.Background(), cfg.URL())
	if err != nil {
		logrus.Fatal("cannot connect to db:", err)
	}
	tx := postgres.New(testDB)
	userRepo := repo.NewUsersRepo(tx)
	notifyService := services.NewNotify()
	userService = users.New(userRepo, notifyService)
}
func TestService_Register(t *testing.T) {
	name := "Aminjon6"
	resp, err := userService.Register(name, "", "@Test123", []string{"+992985068500"}, users.RoleCashier, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, resp.Username, "aminjon6")
	assert.Equal(t, resp.Role, users.RoleCashier)
	assert.NotEmpty(t, resp.ID)
	user, err := userService.SignIn(resp.Username, "@Test123")
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}
func TestService_SignIn(t *testing.T) {
	user, err := userService.SignIn("aminjon", "@Test123")
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}
