package user

import (
	"context"
	model "ecomjc-be/models"
)

type UserService interface {
	GenerateAuthToken(ctx context.Context, userID string) (string, error)
	ParseAuthToken(ctx context.Context, token string) (string, error)
	HashPassword(ctx context.Context, password string) (string, error)
	VerifyEmail(ctx context.Context, authCode string) (bool, error)

	RegisterUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
}
