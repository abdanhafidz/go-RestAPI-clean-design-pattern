package services

import (
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type UserProfileService struct {
	Service[models.AccountDetails, models.AccountDetails]
}

func (s *UserProfileService) Create() {
	userProfile := repositories.CreateAccountDetails(s.Constructor)
	s.Error = userProfile.RowsError
	if userProfile.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no account with given credentials!"
		return
	}
	s.Result = userProfile.Result
}
func (s *UserProfileService) Retrieve() {
	userProfile := repositories.GetAccountDetailsbyId(s.Constructor.AccountID)
	s.Error = userProfile.RowsError
	if userProfile.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no account with given credentials!"
		return
	}
	s.Result = userProfile.Result
}

func (s *UserProfileService) Update() {
	userProfile := repositories.UpdateAccountDetails(s.Constructor)
	s.Error = userProfile.RowsError
	if userProfile.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no account with given credentials!"
		return
	}
	s.Result = userProfile.Result
}
