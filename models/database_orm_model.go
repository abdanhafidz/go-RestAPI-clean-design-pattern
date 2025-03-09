package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Id              uint      `gorm:"primaryKey" json:"id"`
	UUID            uuid.UUID `gorm:"type:uuid" json:"uuid" `
	Email           string    `gorm:"uniqueIndex" json:"email"`
	Password        string    `json:"password"`
	IsEmailVerified bool      `json:"is_email_verified"`
	CreatedAt       time.Time `json:"created_at"`
	DeletedAt       time.Time `json:"deleted_at"`
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
