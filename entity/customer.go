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
