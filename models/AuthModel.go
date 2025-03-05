package models

type AccountData struct {
	IdUser       int
	VerifyStatus string
	ErrVerif     error
}
