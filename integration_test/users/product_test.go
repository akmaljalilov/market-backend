package users

import (
	"context"
	"fmt"
	"market/internal/app/products"
	"market/internal/config"
	repo "market/internal/repository/posgresql"
	postgres "market/internal/repository/posgresql/db"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var productService *products.Service

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
	repo := repo.NewProductsRepo(tx)
	productService = products.NewService(repo)
}
func TestService_RegisterProduct(t *testing.T) {
	list, err := productService.ListMeasurement()
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
	for i, m := range list {
		err = productService.Register(fmt.Sprintf("category_id-%d", i), m.Id)
		assert.NoError(t, err)
	}
}
