package products

type App struct {
	repo Repository
}

// New constructor
func New(repo Repository) *App {
	return &App{repo: repo}
}

func (receiver App) InsertCategory(name string, measurementId int, parentId *int) (int, error) {
	return receiver.repo.InsertCategory(name, measurementId, parentId)
}

func (receiver App) ListMeasurement() ([]Measurement, error) {
	return receiver.repo.ListMeasurement()
}
func (receiver App) InsertProduct(name string, categoryID int) (int, error) {
	return receiver.repo.InsertProduct(name, categoryID)
}
func (receiver App) AddConsumptionPurchaseItem(purchaseItemId int, sum, data string) error {
	return receiver.repo.AddConsumptionPurchaseItem(purchaseItemId, sum, data)
}
func (receiver App) AddPurchaseItem(purchaseId int, productId int, quantity int, sum string, status bool) (int, error) {
	return receiver.repo.AddPurchaseItem(purchaseId, productId, quantity, sum, status)
}
func (receiver App) CreatePurchase(dealerId string) (int, error) {
	return receiver.repo.CreatePurchase(dealerId)
}
func (receiver App) GetProductsBalance() ([]ProductBalance, error) {
	return receiver.repo.GetProductsBalance()
}
