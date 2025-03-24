package services

import (
	"time"

	"godp.abdanhafidz.com/models"
)

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
	if len(constructor) == 1 {
		return &Service[TConstructor, TResult]{}
	}

	return &Service[TConstructor, TResult]{
		Constructor: constructor[0],
	}
}

func CalculateDueTime(duration time.Duration) time.Time {
	return time.Now().Add(duration)
}
