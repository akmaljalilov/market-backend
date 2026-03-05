package products

type Repository interface {
	Register(name string, measurementId int, parentId *int) (int, error)
	InsertProduct(name string, categoryID int) (int, error)
	AddConsumptionPurchaseItem(purchaseItemId int, sum string) error
	AddPurchaseItem(purchaseId int, productId int, quantity int, status bool) (int, error)
	CreatePurchase(dealerId string) (int, error)
	ListMeasurement() ([]Measurement, error)
}
