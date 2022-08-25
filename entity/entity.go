package entity

import "time"

type Order struct {
	IdOrder            int       `gorm:"column:id;primaryKey" json:"id"`
	CustomerIdCustomer int       `gorm:"column:user_id" json:"user_id"`
	Customer           Customer  `gorm:"foreignKey:CustomerIdCustomer" json:"customer"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt           time.Time `gorm:"column:updated_at" json:"update_at"`
}

type OderDetail struct {
	IdOrderDetails   int       `gorm:"column:id" json:"id"`
	OrderIdOrder     int       `gorm:"column:order_id" json:"order_id"`
	Order            Order     `gorm:"foreignKey:OrderIdOrder" json:"order"`
	ProductIdProduct int       `gorm:"column:product_id" json:"id_product"`
	Product          Product   `gorm:"foreignKey:ProductIdProduct" json:"product"`
	Qty              int       `gorm:"column:quantity" json:"quantity"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt         time.Time `gorm:"column:updated_at" json:"update_at"`
}

func (OderDetail) TableName() string {
	return "order_items"
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
