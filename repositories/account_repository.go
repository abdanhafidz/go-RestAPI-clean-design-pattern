package repositories

import (
	"fmt"

	"godp.abdanhafidz.com/models"
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

func GetAccountbyId(account_id uint) Repository[models.Account, models.Account] {
	repo := Construct[models.Account, models.Account](
		models.Account{Id: account_id},
	)
	repo.Transactions(
		WhereGivenConstructor[models.Account, models.Account],
		Find[models.Account, models.Account],
	)
	return *repo
}
func UpdateAccount(account models.Account) Repository[models.Account, models.Account] {
	repo := Construct[models.Account, models.Account](
		account,
	)
	Update(repo)
	return *repo
}
func GetAccountDetailsbyId(account_id uint) Repository[models.AccountDetails, models.AccountDetails] {
	repo := Construct[models.AccountDetails, models.AccountDetails](
		models.AccountDetails{AccountID: account_id},
	)
	fmt.Println("Account ID:", repo.Constructor.AccountID)
	repo.Transactions(
		WhereGivenConstructor[models.AccountDetails, models.AccountDetails],
		Find[models.AccountDetails, models.AccountDetails],
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
func CreateAccountDetails(accountDetails models.AccountDetails) Repository[models.AccountDetails, models.AccountDetails] {
	repo := Construct[models.AccountDetails, models.AccountDetails](
		accountDetails,
	)
	fmt.Println(accountDetails)
	fmt.Println("Account ID : ", accountDetails.AccountID)
	Create(repo)
	return *repo
}
func UpdateAccountDetails(accountDetails models.AccountDetails) Repository[models.AccountDetails, models.AccountDetails] {
	fmt.Println("Account ID : ", accountDetails.AccountID)
	repo := Construct[models.AccountDetails, models.AccountDetails](
		models.AccountDetails{AccountID: accountDetails.AccountID},
	)
	repo.Transaction.Where("account_id = ?", accountDetails.AccountID).First(&repo.Constructor)
	accountDetails.ID = repo.Constructor.ID
	repo.Transaction.Updates(accountDetails)
	repo.Result = accountDetails
	return *repo
}
