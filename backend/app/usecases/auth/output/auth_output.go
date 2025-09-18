package output

import "time"

// LoginUserRes ログインユーザー情報レスポンス
type LoginUserRes struct {
	UserId   string `json:"user_id"`
	LoginId  string `json:"login_id"`
	UserName string `json:"user_name"`
	IsAdmin  bool   `json:"is_admin"`
}

// Login ログインレスポンス
type Login struct {
	SessionId             string       `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	RefreshToken          string       `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  LoginUserRes `json:"user"`
}

// RenewAccessToken アクセストークン更新レスポンス
type RenewAccessToken struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}
