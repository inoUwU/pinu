package repositories

import (
	"context"
	"fmt"
	"github.com/samber/do"
	"github.com/uptrace/bun"
	"inoUwU/pinu/app/domain/entities"
	"inoUwU/pinu/app/domain/repositories"
)

// UserRepositoryImpl ユーザーリポジトリの実装（アダプター）
type UserRepositoryImpl struct {
	db *bun.DB
}

// NewUserRepository ユーザーリポジトリの実装を生成する
func NewUserRepository(i *do.Injector) (repositories.IUserRepository, error) {
	db := do.MustInvokeNamed[*bun.DB](i, "db")
	return &UserRepositoryImpl{db: db}, nil
}

// GetAllUsers 全てのユーザーを取得する
func (r *UserRepositoryImpl) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	users := make([]entities.User, 0)
	if err := r.db.NewSelect().Model(&users).Scan(ctx); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID IDでユーザーを取得する
func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id string) (*entities.User, error) {
	user := new(entities.User)
	if err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, nil // ユーザーが見つからない場合はnilを返す
	}
	fmt.Println(*user)
	return user, nil
}

// CreateUser ユーザーを作成する
func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *entities.User) error {
	const query = `INSERT INTO users (id, name, email) VALUES (?, ?, ?)`

	_, err := r.db.Exec(query, user.ID, user.Name, user.Email)
	return err
}

// UpdateUser ユーザーを更新する
func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user *entities.User) error {
	const query = `UPDATE users SET name = ?, email = ? WHERE id = ?`

	_, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

// DeleteUser ユーザーを削除する
func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	const query = `DELETE FROM users WHERE id = ?`

	_, err := r.db.Exec(query, id)
	return err
}
