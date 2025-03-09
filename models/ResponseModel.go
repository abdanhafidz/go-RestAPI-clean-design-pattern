package models

type AuthenticatedUser struct {
	Account Account `json:"account"`
	Token   string  `json:"token"`
}
