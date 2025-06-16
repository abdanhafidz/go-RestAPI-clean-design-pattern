package services

import (
	"context"
	"errors"

	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
	"gorm.io/gorm"
)

type AuthenticationService interface {
	Service
	Authenticate(ctx context.Context, email string, password string) (res models.AuthenticatedUser)
	Create(ctx context.Context, email string, password string) (res models.AuthenticatedUser)
}

type authenticationService struct {
	*service[repositories.AccountRepository]
	userProfileService UserProfileService
}

func NewAuthenticationService(accountRepository repositories.AccountRepository, userProfileService UserProfileService) AuthenticationService {
	service := authenticationService{}
	service.repository = accountRepository
	service.userProfileService = userProfileService
	return &service
}
func (s *authenticationService) Authenticate(ctx context.Context, email string, password string) (res models.AuthenticatedUser) {
	account := s.repository.GetAccountByEmail(ctx, email)

	if s.ThrowsRepoException() {
		return
	}

	if VerifyPassword(account.Password, password) != nil {
		s.ThrowsException(&s.exception.Unauthorized, "There is no account with given credentials!")
		return
	}

	token, err_tok := GenerateToken(&account)

	if err_tok != nil {
		s.errors = errors.Join(s.errors, err_tok)
	}
	return models.AuthenticatedUser{
		Account: account,
		Token:   token,
	}

}

func validPassword(s *authenticationService, password string) bool {
	if len(password) < 8 {
		s.ThrowsException(&s.exception.BadRequest, "Your password is less than 8 characters!")
		return false
	}
	return true
}
func (s *authenticationService) Create(ctx context.Context, email string, password string) (res models.AuthenticatedUser) {

	if !validPassword(s, password) {
		return
	}

	hashed_password, err_hash := HashPassword(password)

	s.errors = err_hash

	account := s.repository.CreateAccount(ctx, email, hashed_password)

	if errors.Is(s.repository.RowsError(), gorm.ErrDuplicatedKey) {
		s.exception.DataDuplicate = true
		s.exception.Message = "Account with email " + email + " already exists!"
		return
	} else if errors.Is(s.repository.RowsError(), gorm.ErrModelAccessibleFieldsRequired) ||
		errors.Is(s.repository.RowsError(), gorm.ErrInvalidData) ||
		errors.Is(s.repository.RowsError(), gorm.ErrInvalidValue) ||
		errors.Is(s.repository.RowsError(), gorm.ErrInvalidField) {
		s.exception.BadRequest = true
		s.exception.Message = "Bad request!"
		return
	}

	s.userProfileService.Create(ctx, account.Id)
	s.Authenticate(ctx, email, password)
	res.Account = account
	return
}

// func (s *AuthenticationService) ChangePassword(ctx context.Context, oldPassword string, newPassword string) {
// 	if len(newPassword) < 8 {
// 		s.exception.InvalidPasswordLength = true
// 		s.exception.Message = "Password must have at least 8 characters!"
// 		return
// 	}
// 	accountData := repositories.GetAccountbyId(s.constructor.Id)

// 	// if s.Repoexception(accountData) {
// 	// 	return
// 	// }
// 	if VerifyPassword(accountData.Result.Password, oldPassword) != nil {
// 		s.ThrowsException(&s.exception.Unauthorized, "incorrect old password!")
// 		return
// 	}
// 	accountData.Result.Password = newPassword
// 	changePassword := repositories.UpdateAccount(accountData.Result)
// 	changePassword.Result.Password = "SECRET"
// 	s.result = models.AuthenticatedUser{
// 		Account: changePassword.Result,
// 	}
// 	s.errors = changePassword.Rowserror
// }
