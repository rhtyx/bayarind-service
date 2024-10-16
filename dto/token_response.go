package dto

type TokenResponse struct {
	RefreshToken  string `json:"refresh_token"`
	AccessToken   string `json:"access_token"`
	HMACSecretKey string `json:"hmac_secret_key"`
}
