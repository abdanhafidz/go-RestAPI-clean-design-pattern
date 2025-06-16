package factory

import (
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/repositories"
	"godp.abdanhafidz.com/services"
)

func NewAuthenticationModule() controller.AuthController {

	accountRepository := repositories.NewAccountRepository()
	accountDetailRepository := repositories.NewAccountDetailRepository()
	userProfileService := services.NewUserProfileService(accountDetailRepository)
	accountService := services.NewAuthenticationService(accountRepository, userProfileService)
	accountController := controller.NewAuthController(accountService)

	return accountController
}

func NewUserProfileModule() controller.UserController {

	accountDetailRepository := repositories.NewAccountDetailRepository()
	userProfileService := services.NewUserProfileService(accountDetailRepository)
	userController := controller.NewUserController(userProfileService)

	return userController
}
