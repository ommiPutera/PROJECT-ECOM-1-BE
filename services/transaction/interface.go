package transaction

import (
	"context"
	model "ecomjc-be/models"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, Transaction *model.Transaction) (*model.Transaction, error)
	GetTransactiones(ctx context.Context) ([]*model.Transaction, error)
	GetTransactionByID(ctx context.Context, TransactionID int) (*model.Transaction, error)
	UpdateTransaction(ctx context.Context, user *model.Transaction) (*model.Transaction, error)
}
