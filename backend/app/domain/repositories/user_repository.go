package repositories

import "inoUwU/pinu/app/domain/entities"

// IUserRepository ユーザーリポジトリのインターフェース（ポート）
type IUserRepository interface {
	// GetAllUsers 全てのユーザーを取得する
	GetAllUsers() ([]entities.User, error)

	// GetUserByID IDでユーザーを取得する
	GetUserByID(id string) (*entities.User, error)

	// CreateUser ユーザーを作成する
	CreateUser(user *entities.User) error

	// UpdateUser ユーザーを更新する
	UpdateUser(user *entities.User) error

	// DeleteUser ユーザーを削除する
	DeleteUser(id string) error
}
