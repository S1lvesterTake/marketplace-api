//+build wireinject

package main

import (
	"github.com/S1lvesterTake/marketplace-api/application/use_case/create_product"
	"github.com/google/wire"

	repo "github.com/S1lvesterTake/marketplace-api/infrastructure/persistance/db"
	request "github.com/S1lvesterTake/marketplace-api/infrastructure/transport/http"
	"github.com/jinzhu/gorm"
)

func CreateProductHandler(db *gorm.DB) create_product.CreateProductHandler {
	wire.Build(request.NewRequest, repo.NewProductRepository, create_product.NewCreateProductHandler)
	return create_product.CreateProductHandler{}
}
