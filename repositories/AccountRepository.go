package repositories

import "go-dp.abdanhafidz.com/models"

func GetAccountbyUsernamePassword(username string, password string) Repository[models.Account, models.Account] {
	repo := Construct[models.Account, models.Account](
		models.Account{Username: username, Password: password},
	)
	repo.Transactions(
		Find,
	)
	return *repo
}

func CreateAccount(account models.Account) Repository[models.Account, models.Account] {
	repo := Construct[models.Account, models.Account](
		account,
	)
	repo.Transactions(
		Create,
	)
	return *repo
}
