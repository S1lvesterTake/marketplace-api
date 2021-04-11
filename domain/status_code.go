package domain

import "time"

type StatusCode struct {
	ID          string    `gorm:"column:id;not null;index" json:"id"`
	Description string    `gorm:"column:description" json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
