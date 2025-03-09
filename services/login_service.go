package services

import (
	"errors"

	"go-dp.abdanhafidz.com/middleware"
	"go-dp.abdanhafidz.com/models"
	"go-dp.abdanhafidz.com/repositories"
)

type LoginConstructor struct {
	Email    string
	Password string
}

type AuthenticationService struct {
	Service[LoginConstructor, models.AuthenticatedUser]
}

func (s *AuthenticationService) Authenticate() {
	accountData := repositories.GetAccountbyEmail(s.Constructor.Email)
	if accountData.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "there is no account with given credentials!"
		return
	}
	if middleware.VerifyPassword(accountData.Result.Password, s.Constructor.Password) != nil {
		s.Exception.Unauthorized = true
		s.Exception.Message = "incorrect password!"
		return
	}

	token, err_tok := middleware.GenerateToken(&accountData.Result)

	if err_tok != nil {
		s.Error = errors.Join(s.Error, err_tok)
	}

	accountData.Result.Password = "SECRET"
	s.Result = models.AuthenticatedUser{
		Account: accountData.Result,
		Token:   token,
	}
	s.Error = accountData.RowsError
}

// LoginHandler handles user login
