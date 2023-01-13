package client

import "time"

// UserAttributes represents your personal user account that you use to log in to Lemon Squeezy.
type UserAttributes struct {
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Color           string    `json:"color"`
	AvatarURL       string    `json:"avatar_url"`
	HasCustomAvatar bool      `json:"has_custom_avatar"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// UserApiResponse represents a UserAttributes api response
type UserApiResponse = ApiResponseWithoutRelationships[UserAttributes]
