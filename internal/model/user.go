package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	State     string    `json:"state"`
	UserToken UserToken `json:"user_token"`
}

type UserToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	Expiry       time.Time `json:"expiry"`

	Email   string `json:"email"`
	Picture string `json:"picture"`
}
