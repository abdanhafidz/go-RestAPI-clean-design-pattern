package services

import (
	"errors"
	"sync"
	"time"

	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type (
	service[TRepo repositories.Repository] struct {
		sync.Mutex
		repository TRepo
		exception  models.Exception
		errors     error
	}
	Service interface {
		ThrowsException(*bool, string)
		ThrowsError(error)
		Exception() models.Exception
		Error() error
	}
)

func (s *service[TRepo]) ThrowsException(status *bool, message string) {
	s.Lock()
	*status = true
	s.exception.Message = message
	s.Unlock()
}

func (s *service[TRepo]) ThrowsError(err error) {
	s.Lock()
	s.errors = errors.Join(s.errors, err)
	s.Unlock()
}

func (s *service[TRepo]) Exception() models.Exception {
	return s.exception
}
func (s *service[TRepo]) Error() error {
	return s.errors
}
func CalculateDueTime(duration time.Duration) time.Time {
	return time.Now().Add(duration)
}

func (s *service[TRepo]) ThrowsRepoException() bool {
	s.Lock()
	if s.repository.RowsError() != nil {
		s.ThrowsError(s.repository.RowsError())
		return true
	}
	if s.repository.IsNoRecord() {
		s.ThrowsException(&s.exception.DataNotFound, "No record found")
		return true
	}
	s.Unlock()
	return false
}
