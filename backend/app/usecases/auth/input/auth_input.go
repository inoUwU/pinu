package input

import (
	"inoUwU/pinu/app/domain/user"
)

// Login ログイン用の入力データ
type Login struct {
	LoginId   user.LoginID `json:"login_id"`
	Password  string       `json:"password"`
	IP        string       `json:"ip"`
	UserAgent string       `json:"user_agent"`
}

// RenewAccessToken アクセストークン更新用の入力データ
type RenewAccessToken struct {
	RefreshToken string `json:"refresh_token"`
}

type RevokeSession struct {
	SessionId string `json:"session_id"`
}
