package dto

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=1"`
}
