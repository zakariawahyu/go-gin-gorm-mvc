package entity

import "time"

type Product struct {
	ProductID   int       `gorm:"column:id;primaryKey" json:"id"`
	NameProduct string    `gorm:"column:name" json:"name"`
	Price       int       `json:"price"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt    time.Time `gorm:"column:updated_at" json:"update_at"`
}
