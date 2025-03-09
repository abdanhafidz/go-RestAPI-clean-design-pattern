package repositories

import (
	"go-dp.abdanhafidz.com/models"
)

func GetAccountbyEmail(email string) Repository[models.Account, models.Account] {
	repo := Construct[models.Account, models.Account](
		models.Account{Email: email},
	)
	repo.Transactions(
		WhereGivenConstructor[models.Account, models.Account],
		Find[models.Account, models.Account],
	)
	return *repo
}

func CreateAccount(account models.Account) Repository[models.Account, models.Account] {
	repo := Construct[models.Account, models.Account](
		account,
	)
	Create(repo)
	return *repo
}
