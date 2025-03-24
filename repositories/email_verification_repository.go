package repositories

import (
	"time"

	"godp.abdanhafidz.com/models"
)

func CreateEmailVerification(accountId uint, dueTime time.Time, token uint) Repository[models.EmailVerification, models.EmailVerification] {
	repo := Construct[models.EmailVerification, models.EmailVerification](
		models.EmailVerification{
			AccountID: accountId,
			IsExpired: false,
			ExpiredAt: dueTime,
			Token:     token,
		},
	)
	Create(repo)
	return *repo
}

func GetEmailVerification(account_id uint, token uint) Repository[models.EmailVerification, models.EmailVerification] {
	repo := Construct[models.EmailVerification, models.EmailVerification](
		models.EmailVerification{
			AccountID: account_id,
			IsExpired: false,
			Token:     token,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.EmailVerification, models.EmailVerification],
		Find[models.EmailVerification, models.EmailVerification],
	)
	return *repo
}

func DeleteEmailVerification(token uint) Repository[models.EmailVerification, models.EmailVerification] {
	repo := Construct[models.EmailVerification, models.EmailVerification](
		models.EmailVerification{
			Token: token,
		},
	)

	repo.Transactions(
		WhereGivenConstructor[models.EmailVerification, models.EmailVerification],
		Delete[models.EmailVerification],
	)
	return *repo
}
