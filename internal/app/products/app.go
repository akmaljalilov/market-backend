package products

type App struct {
	repo Repository
}

// New constructor
func New(repo Repository) *App {
	return &App{repo: repo}
}

func (receiver App) Register(name string, measurementId int, parentId *int) error {
	_, err := receiver.repo.Register(name, measurementId, parentId)
	return err
}

func (receiver App) ListMeasurement() ([]Measurement, error) {
	return receiver.repo.ListMeasurement()
}
func (receiver App) InsertProduct(name string, categoryID int) (int, error) {
	return receiver.repo.InsertProduct(name, categoryID)
}
func (receiver App) AddConsumptionPurchaseItem(purchaseItemId int, sum string) error {
	return receiver.repo.AddConsumptionPurchaseItem(purchaseItemId, sum)
}
func (receiver App) AddPurchaseItem(purchaseId int, productId int, quantity int, status bool) (int, error) {
	return receiver.repo.AddPurchaseItem(purchaseId, productId, quantity, status)
}
func (receiver App) CreatePurchase(dealerId string) (int, error) {
	return receiver.repo.CreatePurchase(dealerId)
}
