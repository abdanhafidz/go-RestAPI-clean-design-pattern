package services

import (
	"go-dp.abdanhafidz.com/middleware"
	"go-dp.abdanhafidz.com/models"
	"go-dp.abdanhafidz.com/repositories"
)

type LoginService struct {
	Username string
	Password string
}

type AuthenticationService struct {
	Service[LoginService, models.Account]
}

func (s *AuthenticationService) Authenticate() {
	accountData := repositories.GetAccountbyUsernamePassword(s.Constructor.Username, s.Constructor.Password)
	if accountData.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "there is no account with given username!"
		return
	}

	if middleware.VerifyPassword(accountData.Result.Password, s.Constructor.Password) != nil {
		s.Exception.Unauthorized = true
		s.Exception.Message = "incorrect password!"
		return
	}

	s.Result = accountData.Result
	s.Error = accountData.RowsError
}

// LoginHandler handles user login
