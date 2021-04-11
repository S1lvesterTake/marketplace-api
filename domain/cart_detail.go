package domain

type CartDetail struct {
	BaseModelSoftDelete
	CartID           uint64  `gorm:"column:cart_id;not null;index" json:"cart_id"`
	ProductID        uint64  `gorm:"column:product_id;not null;index" json:"product_id"`
	Quantity         int     `gorm:"column:quantity;not null" json:"quantity"`
	CartDetailAmount float64 `gorm:"column:cart_detail_amount;not null" json:"cart_detail_amount"`
	Product          Product `gorm:"foreignKey:ProductID" json:"product"`
}
