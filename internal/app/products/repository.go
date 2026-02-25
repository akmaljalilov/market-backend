package products

type Repository interface {
	ListMeasurement() ([]Measurement, error)
	Register(name string, measurementId int) (int, error)
}
