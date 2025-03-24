package services

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
	"gorm.io/gorm"
)

type RegisterService struct {
	Service[models.Account, models.Account]
}

func (s *RegisterService) Create() {
	if len(s.Constructor.Password) < 8 {
		s.Exception.InvalidPasswordLength = true
		s.Exception.Message = "Password must have at least 8 characters!"
		return
	}
	hashed_password, err_hash := HashPassword(s.Constructor.Password)
	s.Error = err_hash
	s.Constructor.Password = hashed_password
	s.Constructor.UUID = uuid.NewV4()
	accountCreated := repositories.CreateAccount(s.Constructor)
	if errors.Is(accountCreated.RowsError, gorm.ErrDuplicatedKey) {
		s.Exception.DataDuplicate = true
		s.Exception.Message = "Account with email " + s.Constructor.Email + " already exists!"
		return
	} else if errors.Is(accountCreated.RowsError, gorm.ErrModelAccessibleFieldsRequired) || errors.Is(accountCreated.RowsError, gorm.ErrInvalidData) || errors.Is(accountCreated.RowsError, gorm.ErrInvalidValue) || errors.Is(accountCreated.RowsError, gorm.ErrInvalidField) {
		s.Exception.BadRequest = true
		s.Exception.Message = "Bad request!"
		return
	}
	userProfile := UserProfileService{}
	userProfile.Constructor.AccountID = accountCreated.Result.Id
	userProfile.Create()
	if userProfile.Error != nil {
		s.Error = userProfile.Error
		return
	}
	s.Error = accountCreated.RowsError
	s.Result = accountCreated.Result
	s.Result.Password = "SECRET"
}
