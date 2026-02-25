package users

import (
	"context"
	users "market/internal/app/users"
	"market/internal/config"
	repo "market/internal/repository/posgresql"
	postgres "market/internal/repository/posgresql/db"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var service *users.Service

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
	service = users.NewService(userRepo)
}
func TestService_Register(t *testing.T) {
	resp, err := service.Register("Aminjon", "", "@Test123", []string{"+992985068500"}, users.RoleCashier, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, resp.Username, "aminjon")
	assert.Equal(t, resp.Role, users.RoleCashier)
	assert.NotEmpty(t, resp.ID)
}
