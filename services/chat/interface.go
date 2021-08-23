package chat

import (
	"context"
	model "ecomjc-be/models"
)

type ChatService interface {
	CreateChat(ctx context.Context, Chat *model.Chat) (*model.Chat, error)
	GetChats(ctx context.Context) ([]*model.Chat, error)
	GetChatByID(ctx context.Context, ChatID string) (*model.Chat, error)
	DeleteChat(ctx context.Context, ChatID string) (bool, error)
}
