package repositories

import (
	"context"
	"sync"

	"gorm.io/gorm"
)

type Repository interface {
	Transactions(ctx context.Context, act func(ctx context.Context, tx *gorm.DB))
	FindAllPaginate(ctx context.Context, res any)
	Where(ctx context.Context)
	Find(ctx context.Context, res any)
	Create(ctx context.Context)
	Update(ctx context.Context)
	Query(ctx context.Context, res any)
	Delete(ctx context.Context)
	IsNoRecord() bool
	RowsCount() int
	RowsError() error
}
type PaginationConstructor struct {
	limit  int
	offset int
	filter string
}

type CustomQueryConstructor struct {
	sql    string
	values interface{}
}

type repository[TEntity any] struct {
	sync.Mutex
	entity      TEntity
	pagination  PaginationConstructor
	customQuery CustomQueryConstructor
	transaction *gorm.DB
	rowsCount   int
	noRecord    bool
	rowsError   error
}

func (repo *repository[T1]) RowsError() error {
	return repo.rowsError
}
func (repo *repository[T1]) RowsCount() int {
	return repo.rowsCount
}
func (repo *repository[T1]) IsNoRecord() bool {
	repo.Lock()
	repo.noRecord = repo.transaction.RowsAffected == 0
	repo.Unlock()
	return repo.noRecord
}
func (repo *repository[T1]) Transactions(ctx context.Context, act func(ctx context.Context, tx *gorm.DB)) {
	repo.Lock()
	act(ctx, repo.transaction)
	repo.Unlock()
}
func (repo *repository[T1]) Where(ctx context.Context) {

	tx := repo.transaction.Begin()
	tx.WithContext(ctx).Where(&repo.entity)
	repo.Lock()
	repo.rowsCount = int(tx.RowsAffected)
	repo.noRecord = repo.rowsCount == 0
	repo.rowsError = tx.Error
	repo.Unlock()
}
func (repo *repository[T1]) Find(ctx context.Context, res any) {
	repo.Lock()
	tx := repo.transaction.Begin()
	tx.WithContext(ctx).Find(&res)
	if tx.Error != nil {
		tx.Rollback()
	}
	repo.Lock()
	repo.rowsCount = int(tx.RowsAffected)
	repo.noRecord = repo.rowsCount == 0
	repo.rowsError = tx.Error
	repo.Unlock()
}

func (repo *repository[T1]) FindAllPaginate(ctx context.Context, res any) {

	tx := repo.transaction.Begin()
	tx.WithContext(ctx).Limit(repo.pagination.limit).Offset(repo.pagination.offset).Find(&res)
	if tx.Error != nil {
		tx.Rollback()
	}
	repo.Lock()
	repo.rowsCount = int(tx.RowsAffected)
	repo.noRecord = repo.rowsCount == 0
	repo.rowsError = tx.Error
	repo.Unlock()
}

func (repo *repository[T1]) Create(ctx context.Context) {

	tx := repo.transaction.Begin()
	tx.Create(&repo.entity).Find(&repo.entity)
	if tx.Error != nil {
		tx.Rollback()
	}
	repo.Lock()
	repo.rowsCount = int(tx.RowsAffected)
	repo.noRecord = repo.rowsCount == 0
	repo.rowsError = tx.Error
	repo.Unlock()
}

func (repo *repository[T1]) Update(ctx context.Context) {

	tx := repo.transaction.Begin()
	tx.WithContext(ctx).Save(&repo.entity).Find(&repo.entity)
	if tx.Error != nil {
		tx.Rollback()
	}
	repo.Lock()
	repo.rowsCount = int(tx.RowsAffected)
	repo.noRecord = repo.rowsCount == 0
	repo.rowsError = tx.Error
	repo.Unlock()
}

func (repo *repository[T1]) Delete(ctx context.Context) {

	tx := repo.transaction.Begin()
	tx.WithContext(ctx).Delete(&repo.entity)
	if tx.Error != nil {
		tx.Rollback()
	}
	repo.Lock()
	repo.rowsCount = int(tx.RowsAffected)
	repo.noRecord = repo.rowsCount == 0
	repo.rowsError = tx.Error
	repo.Unlock()
}

func (repo *repository[T1]) Query(ctx context.Context, res any) {

	tx := repo.transaction.Begin()
	tx.WithContext(ctx).Model(&repo.entity).Raw(repo.customQuery.sql, repo.customQuery.values).Scan(&res)
	if tx.Error != nil {
		tx.Rollback()
	}
	repo.Lock()
	repo.rowsCount = int(tx.RowsAffected)
	repo.noRecord = repo.rowsCount == 0
	repo.rowsError = tx.Error
	repo.Unlock()
}
