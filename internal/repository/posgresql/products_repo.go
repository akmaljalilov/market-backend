package postgres

import (
	"context"
	users "market/internal/app/products"
	postgres "market/internal/repository/posgresql/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProductsRepo struct {
	db *postgres.Queries
}

func NewProductsRepo(db *postgres.Queries) *ProductsRepo {
	return &ProductsRepo{db: db}
}

func (r ProductsRepo) Register(name string, measurementId int) (int, error) {
	return r.db.InsertCategory(context.Background(), postgres.InsertCategoryParams{
		Name:          name,
		MeasurementID: measurementId,
	})
}

func (r ProductsRepo) InsertProduct(name string, categoryID int) (int, error) {
	return r.db.InsertProduct(context.Background(), postgres.InsertProductParams{
		Name:       name,
		CategoryID: categoryID,
	})
}

func (r ProductsRepo) AddConsumptionPurchase(purchaseId int, consumption string) error {
	var n pgtype.Numeric
	err := n.ScanScientific(consumption)
	if err != nil {
		return err
	}
	return r.db.AddConsumptionPurchase(context.Background(), postgres.AddConsumptionPurchaseParams{
		Consumption: n,
		ID:          purchaseId,
	})
}

func (r ProductsRepo) AddPurchaseItem(purchaseId int, productId int, quantity int, status bool) (int, error) {
	return r.db.AddPurchaseItem(context.Background(), postgres.AddPurchaseItemParams{
		PurchaseOrderID: purchaseId,
		ProductID:       productId,
		Quantity:        quantity,
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

func (r ProductsRepo) ListMeasurement() ([]users.Measurement, error) {
	list, err := r.db.ListMeasurement(context.Background())
	if err != nil {
		return nil, err
	}
	resp := make([]users.Measurement, len(list))
	for i, item := range list {
		resp[i] = users.Measurement{
			Name: item.Name,
			Id:   item.ID,
		}
	}
	return resp, nil
}
