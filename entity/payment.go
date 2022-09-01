package entity

import "time"

type Payment struct {
	PaymentID         int       `gorm:"primaryKey" json:"id"`
	OrderID           int       `json:"order_id"`
	TransactionCode   string    `json:"transaction_code"`
	PaymentType       string    `json:"payment_type"`
	TransactionTime   string    `json:"transaction_time"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
