package products

type Service struct {
	repo Repository
}

// NewService constructor
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (receiver Service) Register(name string, measurementId int, parentId *int) error {
	_, err := receiver.repo.Register(name, measurementId, parentId)
	return err
}

func (receiver Service) ListMeasurement() ([]Measurement, error) {
	return receiver.repo.ListMeasurement()
}
func (receiver Service) InsertProduct(name string, categoryID int) (int, error) {
	return receiver.repo.InsertProduct(name, categoryID)
}
func (receiver Service) AddConsumptionPurchaseItem(purchaseItemId int, sum string) error {
	return receiver.repo.AddConsumptionPurchaseItem(purchaseItemId, sum)
}
func (receiver Service) AddPurchaseItem(purchaseId int, productId int, quantity int, status bool) (int, error) {
	return receiver.repo.AddPurchaseItem(purchaseId, productId, quantity, status)
}
func (receiver Service) CreatePurchase(dealerId string) (int, error) {
	return receiver.repo.CreatePurchase(dealerId)
}
