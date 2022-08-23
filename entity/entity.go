package entity

import "time"

type Customer struct {
	IdCustomer  int       `gorm:"column:id;primaryKey" json:"id"`
	FullName    string    `gorm:"column:full_name" json:"full_name"`
	Username    string    `gorm:"column:username" json:"username"`
	Email       string    `gorm:"column:email" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt    time.Time `gorm:"column:updated_at" json:"update_at"`
}

type Product struct {
	IdProduct   int
	NameProduct string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Order struct {
	IdOrder    int
	IdCustomer int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type OderDetail struct {
	IdOrderDetails int
	IdOrder        int
	IdProduct      int
	Qty            int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

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
