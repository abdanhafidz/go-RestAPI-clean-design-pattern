package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Id                uint       `gorm:"primaryKey" json:"id"`
	UUID              uuid.UUID  `gorm:"type:uuid" json:"uuid" `
	Email             string     `gorm:"uniqueIndex" json:"email"`
	Password          string     `json:"password"`
	IsEmailVerified   bool       `json:"is_email_verified"`
	IsDetailCompleted bool       `json:"is_detail_completed"`
	CreatedAt         time.Time  `json:"created_at"`
	DeletedAt         *time.Time `json:"deleted_at" gorm:"default:null"`
}

type AccountDetails struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	AccountID     uint       `json:"account_id"`
	InitialName   string     `json:"initial_name"`
	FullName      *string    `json:"full_name"`
	DateOfBirth   *time.Time `json:"date_of_birth"`
	PlaceOfBirth  *string    `json:"place_of_birth"`
	Domicile      *string    `json:"domicile"`
	LastJob       *string    `json:"last_job"`
	Gender        *bool      `json:"gender"`
	LastEducation *string    `json:"last_education"`
	MaritalStatus *bool      `json:"marital_status"`
	Avatar        *string    `json:"avatar"`
	PhoneNumber   *uint      `json:"phone_number"`
}

type EmailVerification struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UUID      uuid.UUID `gorm:"type:uuid" json:"uuid" `
	Token     uint      `json:"token"`
	AccountID uint      `json:"account_id"`
	IsExpired bool      `json:"is_expired"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type ExternalAuth struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	OauthID       string `json:"oauth_id"`
	AccountID     uint   `json:"account_id"`
	OauthProvider string `json:"oauth_provider"`
}

type FCM struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	AccountID uint   `json:"account_id"`
	FCMToken  string `json:"fcm_token"`
}

type ForgotPassword struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UUID      uuid.UUID `gorm:"type:uuid" json:"uuid" `
	Token     uint      `json:"token"`
	AccountID uint      `json:"account_id"`
	IsExpired bool      `json:"is_expired"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// Gorm table name settings
func (Account) TableName() string           { return "account" }
func (AccountDetails) TableName() string    { return "account_details" }
func (EmailVerification) TableName() string { return "email_verifications" }
func (ExternalAuth) TableName() string      { return "extern_auth" }
func (ForgotPassword) TableName() string    { return "forgot_password" }
