package models

import (
	"time"
)

type Account struct {
	IDAccount   uint      `gorm:"primaryKey" json:"id_account"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber int       `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type AccountDetails struct {
	IDDetail    uint      `gorm:"primaryKey" json:"id_detail"`
	IDAccount   uint      `json:"id_account"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	Institution string    `json:"institution"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

// Gorm table name settings
func (Account) TableName() string        { return "account" }
func (AccountDetails) TableName() string { return "account_details" }
