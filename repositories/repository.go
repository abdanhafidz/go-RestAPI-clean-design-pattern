package repositories

import (
	"go-dp.abdanhafidz.com/config"
	"gorm.io/gorm"
)

type Repositories interface {
	FindAllPaginate()
	WhereGivenConstructor()
	Find()
	Create()
	Update()
	CustomQuery()
	Delete()
}
type PaginationConstructor struct {
	Limit  int
	Offset int
	Filter string
}

type CustomQueryConstructor struct {
	SQL    string
	Values interface{}
}

type Repository[TConstructor any, TResult any] struct {
	Constructor TConstructor
	Pagination  PaginationConstructor
	CustomQuery CustomQueryConstructor
	Result      TResult
	Transaction *gorm.DB
	RowsCount   int
	NoRecord    bool
	RowsError   error
}

func Construct[TConstructor any, TResult any](constructor ...TConstructor) *Repository[TConstructor, TResult] {
	if len(constructor) == 1 {
		return &Repository[TConstructor, TResult]{
			Constructor: constructor[0],
			Transaction: config.DB,
		}
	}
	return &Repository[TConstructor, TResult]{
		Constructor: constructor[0],
		Transaction: config.DB.Begin(),
	}
}
func (repo *Repository[T1, T2]) Transactions(transactions ...func(*Repository[T1, T2]) *gorm.DB) {
	for _, tx := range transactions {
		repo.Transaction = tx(repo)
		if repo.RowsError != nil {
			return
		}
	}
}
func WhereGivenConstructor[T1 any, T2 any](repo *Repository[T1, T2]) *gorm.DB {
	tx := repo.Transaction.Where(&repo.Constructor)
	repo.RowsCount = int(tx.RowsAffected)
	repo.NoRecord = repo.RowsCount == 0
	repo.RowsError = tx.Error
	return tx
}
func Find[T1 any, T2 any](repo *Repository[T1, T2]) *gorm.DB {
	tx := repo.Transaction.Find(&repo.Result)
	repo.RowsCount = int(tx.RowsAffected)
	repo.NoRecord = repo.RowsCount == 0
	repo.RowsError = tx.Error
	return tx
}

func FinddAllPaginate[T1 any, T2 any](repo *Repository[T1, T2]) *gorm.DB {
	tx := repo.Transaction.Limit(repo.Pagination.Limit).Offset(repo.Pagination.Offset).Find(&repo.Result)
	repo.RowsCount = int(tx.RowsAffected)
	repo.NoRecord = repo.RowsCount == 0
	repo.RowsError = tx.Error
	return tx
}

func Create[T1 any](repo *Repository[T1, T1]) *gorm.DB {
	tx := repo.Transaction.Create(&repo.Constructor)
	repo.RowsCount = int(tx.RowsAffected)
	repo.NoRecord = repo.RowsCount == 0
	repo.RowsError = tx.Error
	repo.Result = repo.Constructor
	return tx
}

func Update[T1 any](repo *Repository[T1, T1]) *gorm.DB {
	tx := repo.Transaction.Save(&repo.Constructor)
	repo.RowsCount = int(tx.RowsAffected)
	repo.NoRecord = repo.RowsCount == 0
	repo.RowsError = tx.Error
	return tx
}

func Delete[T1 any](repo *Repository[T1, T1]) *gorm.DB {
	tx := repo.Transaction.Delete(&repo.Constructor)
	repo.RowsCount = int(tx.RowsAffected)
	repo.NoRecord = repo.RowsCount == 0
	repo.RowsError = tx.Error
	return tx
}

func CustomQuery[T1 any, T2 any](repo *Repository[T1, T2]) *gorm.DB {
	tx := repo.Transaction.Raw(repo.CustomQuery.SQL, repo.CustomQuery.Values).Scan(&repo.Result)
	repo.RowsCount = int(tx.RowsAffected)
	repo.NoRecord = repo.RowsCount == 0
	repo.RowsError = tx.Error
	return tx
}
