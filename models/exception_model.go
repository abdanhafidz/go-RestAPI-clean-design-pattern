package models

type Exception struct {
	Unauthorized          bool   `json:"unauthorized,omitempty"`
	BadRequest            bool   `json:"bad_request,omitempty"`
	DataNotFound          bool   `json:"data_not_found,omitempty"`
	InternalServerError   bool   `json:"internal_server_error,omitempty"`
	DataDuplicate         bool   `json:"data_duplicate,omitempty"`
	QueryError            bool   `json:"query_error,omitempty"`
	InvalidPasswordLength bool   `json:"invalid_password_length,omitempty"`
	Message               string `json:"message,omitempty"`
}
