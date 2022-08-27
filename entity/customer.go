package entity

import "time"

type Customer struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	FullName    string    `gorm:"not null" json:"full_name"`
	Username    string    `gorm:"not null" json:"username"`
	Email       string    `gorm:"not null" json:"email"`
	PhoneNumber string    `gorm:"not null" json:"phone_number"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdateAt    time.Time `gorm:"not null" json:"update_at"`
	Order       []Order   `json:"orders"`
}

type CustomerWithoutOrder struct {
	ID          int       `json:"id"`
	FullName    string    `json:"full_name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func (CustomerWithoutOrder) TableName() string {
	return "customers"
}

type CustomerResponse struct {
	ID          int       `json:"id"`
	FullName    string    `json:"full_name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func (CustomerResponse) TableName() string {
	return "customers"
}
