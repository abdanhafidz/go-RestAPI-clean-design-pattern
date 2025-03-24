package services

import (
	"math/rand/v2"
	"time"

	uuid "github.com/satori/go.uuid"
	"godp.abdanhafidz.com/config"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type EmailVerificationService struct {
	Service[models.EmailVerification, models.EmailVerification]
}

func (s *EmailVerificationService) Create() {
	accountRepo := repositories.GetAccountbyId(s.Constructor.AccountID)
	if accountRepo.NoRecord {
		s.Error = accountRepo.RowsError
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no account data with given credentials!"
		return
	}

	remainingTime := time.Duration(config.EMAIL_VERIFICATION_DURATION) * time.Hour
	dueTime := CalculateDueTime(remainingTime)

	token := uint(rand.IntN(100000))
	s.Constructor.UUID = uuid.NewV4()
	repo := repositories.CreateEmailVerification(s.Constructor.AccountID, dueTime, token)

	s.Error = repo.RowsError
	s.Result = repo.Result
}

func (s *EmailVerificationService) Validate() {
	repo := repositories.GetEmailVerification(s.Constructor.AccountID, s.Constructor.Token)
	s.Error = repo.RowsError
	if repo.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "Invalid token!"
		return
	}
	s.Result = repo.Result
}

func (s *EmailVerificationService) Delete() {
	repo := repositories.DeleteEmailVerification(s.Constructor.Token)
	s.Error = repo.RowsError
	if repo.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "Invalid token!"
		return
	}
	s.Result = repo.Result
}
