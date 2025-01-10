package transaction

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type TxManager struct {
	db *gorm.DB
}

func NewTxManager(db *gorm.DB) *TxManager {
	return &TxManager{db: db}
}

func (tm *TxManager) Execute(ctx context.Context, fn func(*gorm.DB) error) error {
	return tm.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := fn(tx); err != nil {
			return fmt.Errorf("transaction failed: %w", err)
		}
		return nil
	})
}
