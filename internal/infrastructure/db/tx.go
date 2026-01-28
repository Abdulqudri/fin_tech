package db

import (
	"context"
	"gorm.io/gorm"
)

type Tx interface {
	Commit() error
	Rollback() error
	DB() *gorm.DB
}

type GormTx struct {
	tx *gorm.DB
}

func (t *GormTx) Commit() error {
	return t.tx.Commit().Error
}

func (t *GormTx) Rollback() error {
	return t.tx.Rollback().Error
}

func (t *GormTx) DB() *gorm.DB {
	return t.tx
}

func BeginTx(ctx context.Context, db *gorm.DB) (Tx, error) {
	tx := db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &GormTx{tx: tx}, nil
}
