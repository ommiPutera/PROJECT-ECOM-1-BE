package feedback

import (
	"context"
	model "ecomjc-be/models"
)

type FeedbackService interface {
	CreateFeedback(ctx context.Context, Feedback *model.Feedback) (*model.Feedback, error)
	GetFeedbacks(ctx context.Context) ([]*model.Feedback, error)
	GetFeedbackByID(ctx context.Context, FeedbackID string) (*model.Feedback, error)
	UpdateFeedback(ctx context.Context, Feedback *model.Feedback) (*model.Feedback, error)
	DeleteFeedback(ctx context.Context, FeedbackID string) (bool, error)
}
