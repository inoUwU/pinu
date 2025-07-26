package repositories

import (
	"database/sql"
	"github.com/samber/do"
	"inoUwU/pinu/app/domain/entities"
	"inoUwU/pinu/app/domain/repositories"
)

// UserRepositoryImpl ユーザーリポジトリの実装（アダプター）
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository ユーザーリポジトリの実装を生成する
func NewUserRepository(i *do.Injector) (repositories.IUserRepository, error) {
	db := do.MustInvokeNamed[*sql.DB](i, "user")
	return &UserRepositoryImpl{db: db}, nil
}

// GetAllUsers 全てのユーザーを取得する
func (r *UserRepositoryImpl) GetAllUsers() ([]entities.User, error) {
	const query = `SELECT id, name, email FROM users ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID IDでユーザーを取得する
func (r *UserRepositoryImpl) GetUserByID(id string) (*entities.User, error) {
	const query = `SELECT id, name, email FROM users WHERE id = ?`

	var user entities.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // ユーザーが見つからない場合はnilを返す
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser ユーザーを作成する
func (r *UserRepositoryImpl) CreateUser(user *entities.User) error {
	const query = `INSERT INTO users (id, name, email) VALUES (?, ?, ?)`

	_, err := r.db.Exec(query, user.ID, user.Name, user.Email)
	return err
}

// UpdateUser ユーザーを更新する
func (r *UserRepositoryImpl) UpdateUser(user *entities.User) error {
	const query = `UPDATE users SET name = ?, email = ? WHERE id = ?`

	_, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

// DeleteUser ユーザーを削除する
func (r *UserRepositoryImpl) DeleteUser(id string) error {
	const query = `DELETE FROM users WHERE id = ?`

	_, err := r.db.Exec(query, id)
	return err
}
