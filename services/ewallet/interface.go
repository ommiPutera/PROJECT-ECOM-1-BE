package ewallet

import (
	"context"
	model "ecomjc-be/models"
)

type EWalletService interface {
	CreateEWallet(ctx context.Context, EWallet *model.EWallet) (*model.EWallet, error)
	GetEWallets(ctx context.Context) ([]*model.EWallet, error)
	GetEWalletByID(ctx context.Context, EWalletID string) (*model.EWallet, error)
	UpdateEWallet(ctx context.Context, EWallet *model.EWallet) (*model.EWallet, error)
	DeleteEWallet(ctx context.Context, EWalletID string) (bool, error)
}
