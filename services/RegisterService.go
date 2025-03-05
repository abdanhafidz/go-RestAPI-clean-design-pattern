package services

import (
	"errors"

	"go-dp.abdanhafidz.com/models"
	"go-dp.abdanhafidz.com/repositories"
	"gorm.io/gorm"
)

type RegisterService struct {
	Service[models.Account, models.Account]
}

func (s *RegisterService) Create() {
	accountCreated := repositories.CreateAccount(s.Constructor)
	if errors.Is(accountCreated.RowsError, gorm.ErrDuplicatedKey) {
		s.Exception.DataDuplicate = true
		s.Exception.Message = "There is account registered with given data!"
	} else if errors.Is(accountCreated.RowsError, gorm.ErrModelAccessibleFieldsRequired) || errors.Is(accountCreated.RowsError, gorm.ErrInvalidData) || errors.Is(accountCreated.RowsError, gorm.ErrInvalidValue) || errors.Is(accountCreated.RowsError, gorm.ErrInvalidField) {
		s.Exception.BadRequest = true
		s.Exception.Message = "Bad request!"
	}
}
