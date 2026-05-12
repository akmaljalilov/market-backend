package deps

import (
	"market/internal/app/products"
	"market/internal/app/users"
	"market/internal/services"
)

type Repo struct {
	UserRepo    users.Repository
	ProductRepo products.Repository
}

type Service struct {
	Notify services.Notify
}

type Dependencies struct {
	Repo    Repo
	Service Service
}
