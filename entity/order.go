package entity

import "time"

type Order struct {
	ID                  int                   `gorm:"primaryKey" json:"id"`
	CustomerID          int                   `gorm:"not null" json:"customer_id"`
	Customer            CustomerResponse      `json:"-"`
	Status              string                `json:"status"`
	OrderDetailCustomer []OrderDetailCustomer `gorm:"-" json:"order_detail_customer"`
	OrderDetail         []OrderDetail         `json:"order_details"`
	CreatedAt           time.Time             `json:"created_at"`
	UpdateAt            time.Time             `json:"update_at"`
}

type OrderDetailCustomer struct {
	ID  int `json:"id"`
	Qty int `json:"qty"`
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

type OrderWithCustomer struct {
	ID         int              `gorm:"primaryKey" json:"id"`
	CustomerID int              `gorm:"not null" json:"customer_id"`
	Customer   CustomerResponse `json:"customer"`
	Status     string           `json:"status"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdateAt   time.Time        `json:"update_at"`
}

func (OrderWithCustomer) TableName() string {
	return "orders"
}
