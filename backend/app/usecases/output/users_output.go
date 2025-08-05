package output

import "inoUwU/pinu/app/domain/entities"

type UserOutputPort interface {
}

// GetUsersOutput ユーザー一覧取得の出力パラメータ
type GetUsersOutput struct {
	Users []entities.User `json:"users"`
	Count int             `json:"count"`
}

type CreateUserOutput struct {
}

type UpdateUserOutput struct {
}

type DeleteUserOutput struct {
}
