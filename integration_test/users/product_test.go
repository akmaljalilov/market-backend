package users

import (
	"context"
	"fmt"
	"market/internal/app/products"
	"market/internal/app/users"
	"market/internal/config"
	repo "market/internal/repository/posgresql"
	postgres "market/internal/repository/posgresql/db"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var productApp *products.App
var userApp *users.App

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
	prRepo := repo.NewProductsRepo(tx)
	userRepo := repo.NewUsersRepo(tx)
	productApp = products.New(prRepo)
	userApp = users.New(userRepo, nil)
}
func TestService_RegisterProduct(t *testing.T) {
	list, err := productApp.ListMeasurement()
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
	for i, m := range list {
		_, err = productApp.InsertCategory(fmt.Sprintf("category_id-%d", i), m.Id, nil)
		assert.NoError(t, err)
	}
}
func TestProducts(t *testing.T) {
	categoryId, err := productApp.InsertCategory("cement", 1, nil)
	assert.NoError(t, err)
	productId, err := productApp.InsertProduct("Isfara M(200)", categoryId)
	assert.NoError(t, err)
	assert.NotEmpty(t, productId)
	dealerId, err := userApp.RegisterDealer("Abdu", []string{"+992-92-771-28-29"}, "abdu data")
	assert.NoError(t, err)
	t.Run("Purchase", func(t *testing.T) {
		purchaseId, err := productApp.CreatePurchase(dealerId)
		assert.NoError(t, err)
		purchaseItemId, err := productApp.AddPurchaseItem(purchaseId, productId, 1000, "1030", true)
		assert.NoError(t, err)
		err = productApp.AddConsumptionPurchaseItem(purchaseItemId, "50.00", "Roh")
		assert.NoError(t, err)
		purchaseItemId, err = productApp.AddPurchaseItem(purchaseId, productId, 500, "515", true)
		assert.NoError(t, err)
		err = productApp.AddConsumptionPurchaseItem(purchaseItemId, "15.00", "Korgar")
		assert.NoError(t, err)
		balance, err := productApp.GetProductsBalance()
		assert.NoError(t, err)
		assert.NotEmpty(t, balance)
		balanceProduct := balance[len(balance)-1]
		assert.Equal(t, balanceProduct.ProductId, productId)
		assert.Equal(t, balanceProduct.Sum, float64(1610))
		assert.Equal(t, balanceProduct.Quantity, 1500)
	})
}
func TestSales(t *testing.T) {

}
