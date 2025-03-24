package models

type AccountData struct {
	UserID       int
	VerifyStatus string
	ErrVerif     error
}
