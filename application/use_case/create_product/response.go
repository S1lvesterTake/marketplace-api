package create_product

import (
	"time"

	"github.com/S1lvesterTake/marketplace-api/application/helper"
	"github.com/S1lvesterTake/marketplace-api/domain"
)

type CreateProductResponse struct {
	helper.BaseResponse
	Data CreateProductResponseData `json:"data"`
}
type CreateProductResponseData struct {
	ID          uint64    `json:"id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Stock       uint      `json:"stock"`
	Price       int64     `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func NewCreateProductResponse(data domain.Product, message string, success bool) CreateProductResponse {
	return CreateProductResponse{
		BaseResponse: helper.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: CreateProductResponseData{
			ID:          data.BaseModelSoftDelete.ID,
			ProductName: data.ProductName,
			Description: data.Description,
			Stock:       data.Stock,
			Price:       data.Price,
			CreatedAt:   data.BaseModelSoftDelete.CreatedAt,
			UpdatedAt:   data.BaseModelSoftDelete.UpdatedAt,
			DeletedAt:   data.BaseModelSoftDelete.UpdatedAt,
		},
	}
}
