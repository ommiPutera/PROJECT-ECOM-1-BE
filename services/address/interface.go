package address

import (
	"context"
	model "ecomjc-be/models"
)

type AddressService interface {
	CreateAddress(ctx context.Context, address *model.Address) (*model.Address, error)
	GetAddresses(ctx context.Context) ([]*model.Address, error)
	GetAddressByID(ctx context.Context, addressID int) (*model.Address, error)
	UpdateAddress(ctx context.Context, user *model.Address) (*model.Address, error)
}
