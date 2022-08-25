package entity

import "time"

type Order struct {
	ID         int              `gorm:"primaryKey" json:"id"`
	CustomerID int              `gorm:"not null" json:"customer_id"`
	Customer   CustomerResponse `json:"customer"`
	Status     string           `json:"status"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdateAt   time.Time        `json:"update_at"`
}

type OrderResponse struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
}

func (OrderResponse) TableName() string {
	return "orders"
}
