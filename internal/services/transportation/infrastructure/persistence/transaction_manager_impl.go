package transportationrepositoryimpl

import (
	"context"
	"database/sql"

	transportationrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/repository"
)

type txKeyType struct{}

var txKey = txKeyType{}

type transactionManager struct {
	db *sql.DB
}

// WithTransaction implements authrepository.TransactionManager.
func (tm *transactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	// Inject transaction into context
	txCtx := context.WithValue(ctx, txKey, tx)

	if err := fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func NewTransactionManager(db *sql.DB) transportationrepository.TransactionManager {
	return &transactionManager{db: db}
}
