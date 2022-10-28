package models

import "time"

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url"`
	UserId         uint   `json:"user_id"`
	User           *User
}

type CreateSocialMedia struct {
	Name           string `json:"name" form:"name" valid:"required~name is required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required~socialMediaUrl is required"`
}

type UpdateSocialMedia struct {
	Name           string `json:"name" form:"name" valid:"required~name is required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required~socialMediaUrl is required"`
}

type SocialMediaResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserId         uint       `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
	CreatedAt      *time.Time `json:"created_at"`
	User           *UserSocialMediaResponse
}

type CreateSocialMediaResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserId         uint       `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type UpdateSocialMediaResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserId         uint       `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
