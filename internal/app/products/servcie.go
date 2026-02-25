package products

type Service struct {
	repo Repository
}

// NewService constructor
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (receiver Service) Register(name string, measurementId int) error {
	_, err := receiver.repo.Register(name, measurementId)
	return err
}

func (receiver Service) ListMeasurement() ([]Measurement, error) {
	return receiver.repo.ListMeasurement()
}
