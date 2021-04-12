package create_product

import (
	"github.com/S1lvesterTake/marketplace-api/domain"
	"gopkg.in/go-playground/validator.v9"
)

type CreateProductRequest struct {
	Data struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
		Stock       uint   `json:"stock"`
		Price       int64  `json:"price"`
	}
}

func ValidateRequest(req *CreateProductRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

//RequestMapper mapper request to domain
func RequestMapper(req CreateProductRequest) domain.Product {
	return domain.Product{
		ProductName: req.Data.Name,
		Description: req.Data.Description,
		Stock:       req.Data.Stock,
		Price:       req.Data.Price,
	}
}
