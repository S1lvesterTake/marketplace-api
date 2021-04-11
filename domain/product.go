package domain

type Product struct {
	BaseModelSoftDelete
	ProductName string `gorm:"column:product_name;not null;index" json:"product_name"`
	Description string `gorm:"column:description" json:"description"`
	Stock       uint   `gorm:"column:stock;not null" json:"stock"`
	Price       int64  `gorm:"column:price;not null" json:"price"`
}
