package db

import (
	"context"
	"errors"
	"log"

	"github.com/S1lvesterTake/marketplace-api/application/infrastructure"
	"github.com/S1lvesterTake/marketplace-api/domain"
	"github.com/jinzhu/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) infrastructure.ProductRepository {
	return &productRepository{
		DB: DB,
	}
}

func (c *productRepository) CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	if err := c.DB.Create(&product).Error; err != nil {
		log.Printf("[ProductRepository-CreateProduct] error %s", err)
		return product, err
	}

	return product, nil
}

func (c *productRepository) GetProductByID(ctx context.Context, productID string) (domain.Product, error) {
	product := domain.Product{}
	if c.DB.First(&product, "id = ?", productID).RecordNotFound() {
		log.Printf("[ProductRepository-GetProductByID] Produk dengan id %s tidak ditemukan", productID)
		return product, errors.New("Produk dengan id " + productID + " tidak ditemukan")
	}

	return product, nil
}
