package repositories

import (
	"context"
	"inoUwU/pinu/app/domain/entities"
)

// IUserRepository ユーザーリポジトリのインターフェース（ポート）
type IUserRepository interface {
	// GetAllUsers 全てのユーザーを取得する
	GetAllUsers(ctx context.Context) ([]entities.User, error)

	// GetUserByID IDでユーザーを取得する
	GetUserByID(ctx context.Context, id string) (*entities.User, error)

	// CreateUser ユーザーを作成する
	CreateUser(ctx context.Context, user *entities.User) error

	// UpdateUser ユーザーを更新する
	UpdateUser(ctx context.Context, user *entities.User) error

	// DeleteUser ユーザーを削除する
	DeleteUser(ctx context.Context, id string) error
}
