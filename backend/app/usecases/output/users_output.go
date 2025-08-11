package output

import (
	"inoUwU/pinu/app/domain/entities"
	"time"
)

type UserOutputPort interface {
}

// GetUsersOutput ユーザー一覧取得の出力パラメータ
type GetUsersOutput struct {
	Users []entities.User `json:"users"`
	Count int             `json:"count"`
}

type CreateUserOutput struct {
	User      entities.User    `json:"user"`
	LoginId   entities.LoginID `json:"login_id"`
	PassWord  string           `json:"password"`
	UserId    entities.UserID  `json:"user_id"`
	Name      string           `json:"name"`
	IsAdmin   bool             `json:"is_admin"`
	CreatedAt time.Time        `json:"created_at"`
}

type UpdateUserOutput struct {
	User entities.User `json:"user"`
}

type DeleteUserOutput struct {
	User entities.User `json:"user"`
}
