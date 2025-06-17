package repositories

import (
	"context"

	models "godp.abdanhafidz.com/models"
)

type AccountDetailRepository interface {
	Repository
	CreateAccountDetail(ctx context.Context, account_id uint) (res models.AccountDetails)
	UpdateAccountDetail(ctx context.Context, account_id uint, account_detail models.AccountDetails) (res models.AccountDetails)
	GetByAccountId(ctx context.Context, account_id uint) (res models.AccountDetails)
}

type accountDetailRepository struct {
	*repository[models.AccountDetails]
}

func NewAccountDetailRepository() AccountDetailRepository {
	return &accountDetailRepository{
		repository: &repository[models.AccountDetails]{
			entity: models.AccountDetails{},
		},
	}
}
func (r *accountDetailRepository) CreateAccountDetail(ctx context.Context, account_id uint) (res models.AccountDetails) {
	r.Lock()
	r.entity.AccountID = account_id
	r.Create(ctx)
	r.Unlock()
	return r.entity

}

func (r *accountDetailRepository) UpdateAccountDetail(ctx context.Context, account_id uint, account_detail models.AccountDetails) (res models.AccountDetails) {
	r.Lock()
	r.entity.AccountID = account_id
	r.Where(ctx)
	r.entity = account_detail
	r.Update(ctx)
	r.Unlock()
	return r.entity
}

func (r *accountDetailRepository) GetByAccountId(ctx context.Context, account_id uint) (res models.AccountDetails) {
	r.entity.AccountID = account_id
	r.Where(ctx)
	return r.entity
}
