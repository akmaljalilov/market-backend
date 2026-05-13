package products

type Repository interface {
	InsertCategory(name string, measurementId int, parentId *int) (int, error)
	InsertProduct(name string, categoryID int) (int, error)
	AddConsumptionPurchaseItem(purchaseItemId int, sum, data string) error
	AddPurchaseItem(purchaseId int, productId int, quantity int, sum string, status bool) (int, error)
	CreatePurchase(dealerId string) (int, error)
	GetProductsBalance() ([]ProductBalance, error)
	ListMeasurement() ([]Measurement, error)
}
