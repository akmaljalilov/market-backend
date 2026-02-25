package products

import (
	"time"
)

type Product struct {
	ID string

	CreatedAt time.Time
	UpdatedAt time.Time
}
type Measurement struct {
	Name string
	Id   int
}
