package services

import (
	"context"

	models "godp.abdanhafidz.com/models"
	repositories "godp.abdanhafidz.com/repositories"
)

type UserProfileService interface {
	Service
	Create(ctx context.Context, account_id uint) (res models.AccountDetails)
	Retrieve(ctx context.Context, account_id uint) (res models.AccountDetails)
	Update(ctx context.Context, account_id uint, account_detail models.AccountDetails) (res models.AccountDetails)
}

type userProfileService struct {
	*service[repositories.AccountDetailRepository]
}

func NewUserProfileService(accountDetailRepository repositories.AccountDetailRepository) UserProfileService {
	return &userProfileService{
		service: &service[repositories.AccountDetailRepository]{
			repository: accountDetailRepository,
		},
	}
}
func (s *userProfileService) Create(ctx context.Context, account_id uint) (res models.AccountDetails) {
	accountDetail := s.repository.CreateAccountDetail(ctx, account_id)
	if s.ThrowsRepoException() {
		return
	}
	return accountDetail
}
func (s *userProfileService) Retrieve(ctx context.Context, account_id uint) (res models.AccountDetails) {
	accountDetail := s.repository.GetByAccountId(ctx, account_id)
	if s.ThrowsRepoException() {
		return
	}
	return accountDetail
}

func (s *userProfileService) Update(ctx context.Context, account_id uint, account_detail models.AccountDetails) (res models.AccountDetails) {
	accountDetail := s.repository.UpdateAccountDetail(ctx, account_id, account_detail)
	if s.ThrowsRepoException() {
		return
	}
	return accountDetail
}
