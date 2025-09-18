package output

import (
	"inoUwU/pinu/app/domain/user"
	"time"
)

type UserOutputPort interface {
}

// GetUsersOutput ユーザー一覧取得の出力パラメータ
type GetUsersOutput struct {
	Users []user.User `json:"users"`
}

type CreateUserOutput struct {
	User      user.User    `json:"user"`
	LoginId   user.LoginID `json:"login_id"`
	PassWord  string       `json:"password"`
	UserId    user.UserID  `json:"user_id"`
	Name      string       `json:"name"`
	IsAdmin   bool         `json:"is_admin"`
	CreatedAt time.Time    `json:"created_at"`
}

type UpdateUserOutput struct {
	User user.User `json:"user"`
}

type DeleteUserOutput struct {
	User user.User `json:"user"`
}
