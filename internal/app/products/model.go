package products

import (
	"time"
)

type ProductBalance struct {
	ProductId int
	Sum       float64
	Quantity  int64
}
type Product struct {
	ID string

	CreatedAt time.Time
	UpdatedAt time.Time
}
type Measurement struct {
	Name string
	Id   int
}
