package models

type Exception struct {
	Unauthorized        bool
	BadRequest          bool
	DataNotFound        bool
	InternalServerError bool
	DataDuplicate       bool
	QueryError          bool
	Message             string
}
