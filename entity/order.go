package entity

import "time"

type Order struct {
	ID                  int                       `gorm:"primaryKey" json:"id"`
	CustomerID          int                       `gorm:"not null" json:"customer_id"`
	Customer            Customer                  `json:"-"`
	Status              string                    `json:"status"`
	OrderDetailCustomer []OrderDetailCustomer     `gorm:"-" json:"order_detail_customer"`
	OrderDetail         []OrderDetailWithoutOrder `json:"order_details"`
	CreatedAt           time.Time                 `json:"created_at"`
	UpdateAt            time.Time                 `json:"update_at"`
}

type OrderDetailCustomer struct {
	ID  int `json:"id"`
	Qty int `json:"qty"`
}
