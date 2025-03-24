package services

import (
	"errors"
	"fmt"

	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type AuthenticationService struct {
	Service[models.Account, models.AuthenticatedUser]
}

func (s *AuthenticationService) Authenticate() {
	accountData := repositories.GetAccountbyEmail(s.Constructor.Email)
	if accountData.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "there is no account with given credentials!"
		return
	}
	if VerifyPassword(accountData.Result.Password, s.Constructor.Password) != nil {
		s.Exception.Unauthorized = true
		s.Exception.Message = "incorrect password!"
		return
	}

	token, err_tok := GenerateToken(&accountData.Result)

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

func (s *AuthenticationService) Update(oldPassword string, newPassword string) {
	if len(newPassword) < 8 {
		s.Exception.InvalidPasswordLength = true
		s.Exception.Message = "Password must have at least 8 characters!"
		return
	}
	accountData := repositories.GetAccountbyId(s.Constructor.Id)

	if accountData.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "there is no account with given credentials!"
		return
	}
	fmt.Println("Result Password", accountData.Result.Password)
	fmt.Println("old Password given", oldPassword)
	if VerifyPassword(accountData.Result.Password, oldPassword) != nil {
		s.Exception.Unauthorized = true
		s.Exception.Message = "incorrect old password!"
		return
	}
	accountData.Result.Password = newPassword
	changePassword := repositories.UpdateAccount(accountData.Result)
	changePassword.Result.Password = "SECRET"
	s.Result = models.AuthenticatedUser{
		Account: changePassword.Result,
	}
	s.Error = changePassword.RowsError
}

// LoginHandler handles user login
