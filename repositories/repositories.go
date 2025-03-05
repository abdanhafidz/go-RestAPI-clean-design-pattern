package repositories

import (
	"gorm.io/gorm"
)

type Repositories interface {
	FindAllPaginate()
	Where()
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
	if len(constructor) == 0 {
		return &Repository[TConstructor, TResult]{
			Transaction: db.Begin(),
		}
	}
	return &Repository[TConstructor, TResult]{
		Constructor: constructor[0],
		Transaction: db.Begin(),
	}
}
func (repo *Repository[T1, T2]) Transactions(transactions ...func(*Repository[T1, T2])) {
	i := 1
	for _, tx := range transactions {
		tx(repo)
		repo.RowsError = repo.Transaction.Error
		if repo.RowsError != nil {
			repo.Transaction.Rollback()
			return
		} else {
			repo.Transaction.SavePoint("Save Point : " + string(i))
			repo.Transaction.Commit()
		}
	}
}
func WhereGivenConstructor[T1 any, T2 any](repo *Repository[T1, T2]) {
	repo.Transaction.Where(&repo.Constructor)
}
func Find[T1 any, T2 any](repo *Repository[T1, T2]) {
	repo.Transaction.Find(&repo.Result)
}

func FinddAllPaginate[T1 any, T2 any](repo *Repository[T1, T2]) {
	repo.Transaction.Limit(repo.Pagination.Limit).Offset(repo.Pagination.Offset).Find(&repo.Result)
}

func Create[T1 any](repo *Repository[T1, T1]) {
	repo.Transaction.Create(&repo.Constructor)
	repo.Result = repo.Constructor
}

func Update[T1 any](repo *Repository[T1, T1]) {
	repo.Transaction.Save(&repo.Constructor)
	repo.Result = repo.Constructor
}

func Delete[T1 any](repo *Repository[T1, T1]) {
	repo.Transaction.Delete(&repo.Constructor)
}

func CustomQuery[T1 any, T2 any](repo *Repository[T1, T2]) {
	repo.Transaction.Raw(repo.CustomQuery.SQL, repo.CustomQuery.Values).Scan(&repo.Result)
}
