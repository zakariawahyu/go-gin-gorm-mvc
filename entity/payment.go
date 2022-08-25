package entity

import "time"

type Payment struct {
	IdPayment         int
	IdOrder           int
	IdTransaction     string
	PaymentType       string
	GrossAmount       string
	TransactionTIme   string
	TransactionStatus string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
