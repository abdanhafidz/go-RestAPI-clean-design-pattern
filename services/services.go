package services

import "go-dp.abdanhafidz.com/models"

type (
	Services interface {
		Retrieve()
		Update()
		Create()
		Delete()
		Validate()
		Authenticate()
		Authorize()
	}
	Service[TConstructor any, TResult any] struct {
		Constructor TConstructor
		Result      TResult
		Exception   models.Exception
		Error       error
	}
)

func Construct[TConstructor any, TResult any](constructor ...TConstructor) *Service[TConstructor, TResult] {
	if len(constructor) == 0 {
		return &Service[TConstructor, TResult]{}
	}

	return &Service[TConstructor, TResult]{
		Constructor: constructor[0],
	}
}
