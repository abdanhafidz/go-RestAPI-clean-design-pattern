package models

type SuccessResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
	MetaData any    `json:"meta_data"`
}

type ErrorResponse struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Errors   Exception `json:"errors"`
	MetaData any       `json:"meta_data"`
}
type AuthenticatedUser struct {
	Account Account `json:"account"`
	Token   string  `json:"token"`
}
