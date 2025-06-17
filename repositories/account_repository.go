package repositories

import (
	"context"

	models "godp.abdanhafidz.com/models"
)

type AccountRepository interface {
	Repository
	CreateAccount(ctx context.Context, email string, password string) (res models.Account)
	GetAccountByEmail(ctx context.Context, email string) (res models.Account)
}
type accountRepository struct {
	*repository[models.Account]
}

func NewAccountRepository() AccountRepository {
	return &accountRepository{
		repository: &repository[models.Account]{
			entity: models.Account{},
		},
	}
}
func (r *accountRepository) CreateAccount(ctx context.Context, email string, password string) (res models.Account) {
	r.entity.Email = email
	r.entity.Password = password
	r.Create(ctx)
	return r.entity
}

func (r *accountRepository) GetAccountByEmail(ctx context.Context, email string) (res models.Account) {
	r.entity.Email = email
	r.Find(ctx, res)
	return res
}
