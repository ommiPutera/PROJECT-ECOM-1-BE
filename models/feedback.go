package model

import "time"

type Feedback struct {
	ID             int       `json:"id"`
	UserID         string    `json:"user_id"`
	ReportedUserID string    `json:"reported_user_id"`
	ProductID      int       `json:"product_id"`
	Rating         uint8     `json:"rating"`
	Message        string    `json:"message"`
	Image          string    `json:"image"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsEdited       bool      `json:"is_edited"`
	IsDeleted      bool      `json:"is_deleted"`
}
