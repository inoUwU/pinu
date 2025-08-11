package repositories

import (
	"context"
	"fmt"
	"inoUwU/pinu/app/domain/entities"
	"inoUwU/pinu/app/domain/repositories"
	"inoUwU/pinu/app/infrastructure/models"

	"github.com/samber/do"
	"github.com/uptrace/bun"
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
	tmp_users := make([]models.User, 0)
	if err := r.db.NewSelect().Model(&tmp_users).Scan(ctx); err != nil {
		return nil, err
	}

	users := make([]entities.User, len(tmp_users))
	for i := 0; i < len(tmp_users); i++ {
		model := tmp_users[i]
		users[i] = entities.User{
			USER_ID:       entities.UserID(model.USER_ID),
			LOGIN_ID:      entities.LoginID(model.LOGIN_ID),
			PASSWORD_HASH: model.PASSWORD_HASH,
			NAME:          model.NAME,
			IS_ADMIN:      model.IS_ADMIN,
			CREATED_AT:    model.CREATED_AT,
		}
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
	modelUser := &models.User{
		USER_ID:       string(user.USER_ID),
		LOGIN_ID:      string(user.LOGIN_ID),
		PASSWORD_HASH: user.PASSWORD_HASH,
		PASSWORD_SALT: user.PASSWORD_SALT,
		NAME:          user.NAME,
		IS_ADMIN:      user.IS_ADMIN,
		CREATED_AT:    user.CREATED_AT,
	}

	_, err := r.db.NewInsert().Model(modelUser).Exec(ctx)
	return err
}

// UpdateUser ユーザーを更新する
func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user *entities.User) error {
	modelUser := &models.User{
		USER_ID:       string(user.USER_ID),
		LOGIN_ID:      string(user.LOGIN_ID),
		PASSWORD_HASH: user.PASSWORD_HASH,
		PASSWORD_SALT: user.PASSWORD_SALT,
		NAME:          user.NAME,
		IS_ADMIN:      user.IS_ADMIN,
		CREATED_AT:    user.CREATED_AT,
	}
	_, err := r.db.NewUpdate().Model(modelUser).WherePK().Exec(ctx)
	return err
}

// DeleteUser ユーザーを削除する
func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	_, err := r.db.NewDelete().Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
