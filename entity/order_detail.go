package entity

import "time"

type OrderDetail struct {
	OrderID   int               `gorm:"not null" json:"order_id"`
	Order     OrderWithCustomer `json:"order"`
	ProductID int               `gorm:"not null" json:"product_id"`
	Product   Product           `json:"product"`
	Qty       int               `json:"qty"`
	Amount    int               `json:"amount"`
	CreatedAt time.Time         `json:"created_at"`
	UpdateAt  time.Time         `json:"update_at"`
}
