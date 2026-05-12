package deps

import (
	"context"
	"market/internal/config"
	repo "market/internal/repository/posgresql"
	postgres "market/internal/repository/posgresql/db"
	"market/internal/services"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func Init(appCfg *config.Config) *Dependencies {
	testDB, err := pgxpool.New(context.Background(), appCfg.Postgres.URL())
	if err != nil {
		logrus.Fatal("cannot connect to db:", err)
	}
	tx := postgres.New(testDB)
	userRepo := repo.NewUsersRepo(tx)
	productRepo := repo.NewProductsRepo(tx)
	notifyService := services.NewNotify()
	return &Dependencies{
		Repo: Repo{
			UserRepo:    userRepo,
			ProductRepo: productRepo,
		},
		Service: Service{
			Notify: notifyService,
		},
	}
}
