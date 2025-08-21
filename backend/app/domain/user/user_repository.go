package user

import (
	"context"
)

// IUserRepository ユーザーリポジトリのインターフェース（ポート）
type IUserRepository interface {
	// 全てのユーザーを取得する
	GetAllUsers(ctx context.Context) ([]User, error)

	// IDでユーザーを取得する
	GetUserByID(ctx context.Context, id string) (*User, error)

	// ユーザーを作成する
	CreateUser(ctx context.Context, user *User) error

	// ユーザーを更新する
	UpdateUser(ctx context.Context, user *User) error

	// ユーザーを削除する
	DeleteUser(ctx context.Context, id string) error
}
