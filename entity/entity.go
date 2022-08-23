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
	IdProduct   int       `gorm:"column:id;primaryKey" json:"id"`
	NameProduct string    `gorm:"column:name" json:"product_name"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt    time.Time `gorm:"column:updated_at" json:"update_at"`
}

type Order struct {
	IdOrder            int       `gorm:"column:id;primaryKey" json:"id"`
	CustomerIdCustomer int       `gorm:"column:user_id" json:"user_id"`
	Status             string    `json:"status"`
	Customer           Customer  `gorm:"foreignKey:CustomerIdCustomer" json:"customer"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt           time.Time `gorm:"column:updated_at" json:"update_at"`
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
