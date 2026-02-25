package deps

import (
	"context"
	"market/internal/config"
	repo "market/internal/repository/posgresql"
	postgres "market/internal/repository/posgresql/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func Init(appCfg *config.Config) *Services {
	testDB, err := pgxpool.New(context.Background(), appCfg.Postgres.URL())
	if err != nil {
		logrus.Fatal("cannot connect to db:", err)
	}
	tx := postgres.New(testDB)
	userRepo := repo.NewUsersRepo(tx)
	return &Services{
		UserRepo: userRepo,
	}
}
