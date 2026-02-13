package user

import (
	"context"
)

// UserStore ユーザー永続化のポート
type UserStore interface {
	// GetAllUsers 全てのユーザーを取得する
	GetAllUsers(ctx context.Context) ([]User, error)

	// GetUserByID IDでユーザーを取得する
	GetUserByID(ctx context.Context, id string) (*User, error)

	// GetUserByLoginID ログインIDでユーザーを取得する
	GetUserByLoginID(ctx context.Context, loginID string) (*User, error)

	// CreateUser ユーザーを作成する
	CreateUser(ctx context.Context, user *User) error

	// UpdateUser ユーザーを更新する
	UpdateUser(ctx context.Context, user *User) error

	// DeleteUser ユーザーを削除する
	DeleteUser(ctx context.Context, id string) error
}
