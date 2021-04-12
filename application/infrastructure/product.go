package infrastructure

import (
	"context"

	"github.com/S1lvesterTake/marketplace-api/domain"
)

//ProductRepository repository
type ProductRepository interface {
	CreateProduct(context.Context, domain.Product) (domain.Product, error)
	GetProductByID(context.Context, string) (domain.Product, error)
}
