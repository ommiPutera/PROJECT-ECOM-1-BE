package model

import "time"

type Shop struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Domain         string    `json:"domain"`
	PhoneNumber    string    `json:"phone_number"`
	ProfilePicture string    `json:"profile_picture"`
	Followers      int       `json:"followers"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsVerified     bool      `json:"is_verified"`
	IsDeleted      bool      `json:"is_deleted"`
}
