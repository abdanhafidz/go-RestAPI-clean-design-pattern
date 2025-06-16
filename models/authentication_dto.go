package models

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Phone    int    `json:"phone"`
	Password string `json:"password" binding:"required"`
}

type CreateEmailVerificationRequest struct {
	AccountID int `json:"account_id" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required" `
	NewPassword string `json:"new_password" binding:"required" `
}

type AuthenticatedUser struct {
	Account Account `json:"account"`
	Token   string  `json:"token"`
}
