package postgres

import (
	"context"
	"market/internal/app/products"
	postgres "market/internal/repository/posgresql/db"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProductsRepo struct {
	db *postgres.Queries
}

func NewProductsRepo(db *postgres.Queries) *ProductsRepo {
	return &ProductsRepo{db: db}
}

func (r ProductsRepo) InsertCategory(name string, measurementId int, parentId *int) (int, error) {
	return r.db.InsertCategory(context.Background(), postgres.InsertCategoryParams{
		Name:          name,
		MeasurementID: measurementId,
		CategoryID:    parentId,
	})
}

func (r ProductsRepo) InsertProduct(name string, categoryID int) (int, error) {
	return r.db.InsertProduct(context.Background(), postgres.InsertProductParams{
		Name:       name,
		CategoryID: categoryID,
	})
}

func (r ProductsRepo) AddConsumptionPurchaseItem(purchaseItemId int, sum, data string) error {
	var n pgtype.Numeric
	err := n.ScanScientific(sum)
	if err != nil {
		return err
	}
	return r.db.AddExpensesPurchaseItem(context.Background(), postgres.AddExpensesPurchaseItemParams{
		Sum:            n,
		PurchaseItemID: purchaseItemId,
		Data:           &data,
	})
}

func (r ProductsRepo) AddPurchaseItem(purchaseId int, productId int, quantity int, sum string, status bool) (int, error) {
	var n pgtype.Numeric
	err := n.ScanScientific(sum)
	if err != nil {
		return -1, err
	}
	return r.db.AddPurchaseItem(context.Background(), postgres.AddPurchaseItemParams{
		PurchaseOrderID: purchaseId,
		ProductID:       productId,
		Quantity:        quantity,
		Price:           n,
		Status:          status,
	})
}

func (r ProductsRepo) CreatePurchase(dealerId string) (int, error) {
	id, err := uuid.Parse(dealerId)
	if err != nil {
		return 0, err
	}
	return r.db.CreatePurchase(context.Background(), pgtype.UUID{
		Bytes: id,
		Valid: true,
	})
}

func (r ProductsRepo) ListMeasurement() ([]products.Measurement, error) {
	list, err := r.db.ListMeasurement(context.Background())
	if err != nil {
		return nil, err
	}
	resp := make([]products.Measurement, len(list))
	for i, item := range list {
		resp[i] = products.Measurement{
			Name: item.Name,
			Id:   item.ID,
		}
	}
	return resp, nil
}

func (r ProductsRepo) GetProductsBalance() ([]products.ProductBalance, error) {
	list, err := r.db.GetProductsBalance(context.Background())
	if err != nil {
		return nil, err
	}
	resp := make([]products.ProductBalance, len(list))
	for i, item := range list {
		val, err := item.Sum.Float64Value()
		if err != nil {
			log.Error(err)
			continue
		}
		quantity := 0
		if item.Quantity != nil {
			quantity = *item.Quantity
		}
		resp[i] = products.ProductBalance{
			ProductId: item.ProductID,
			Sum:       val.Float64,
			Quantity:  quantity,
		}
	}
	return resp, nil
}
