package repositories

import (
	"context"
	"inoUwU/pinu/app/domain/user"
	"inoUwU/pinu/app/infrastructure/models"

	"inoUwU/pinu/app/infrastructure/ctx"

	"github.com/samber/do"
	"github.com/uptrace/bun"
)

// UserRepositoryImpl ユーザーリポジトリの実装（アダプター）
type UserRepositoryImpl struct {
	db *bun.DB
}

// NewUserRepository ユーザーリポジトリの実装を生成する
func NewUserRepository(i *do.Injector) (user.UserRepository, error) {
	db := do.MustInvokeNamed[*bun.DB](i, "db")
	return &UserRepositoryImpl{db: db}, nil
}

// GetAllUsers 全てのユーザーを取得する
func (r *UserRepositoryImpl) GetAllUsers(ctx context.Context) ([]user.User, error) {
	tmpUsers := make([]models.UserModel, 0)
	if err := r.db.NewSelect().Model(&tmpUsers).Scan(ctx); err != nil {
		return nil, err
	}

	users := make([]user.User, len(tmpUsers))
	for i := 0; i < len(tmpUsers); i++ {
		model := tmpUsers[i]
		users[i] = user.User{
			UserID:       user.UserID(model.UserID),
			LoginID:      user.LoginID(model.LoginID),
			PasswordHash: model.PasswordHash,
			Name:         model.Name,
			IsAdmin:      model.IsAdmin,
			CreatedAt:    model.CreatedAt,
		}
	}

	return users, nil
}

// GetUserByLoginID ログインIDでユーザーを取得する
func (r *UserRepositoryImpl) GetUserByLoginID(ctx context.Context, loginID string) (*user.User, error) {
	model := new(models.UserModel)

	var selector *bun.SelectQuery

	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = r.db.NewSelect()
	}

	selector.Model(model).Where("login_id = ?", loginID)
	if err := selector.Scan(ctx); err != nil {
		return nil, err
	}

	return mapUserModelToDomain(model), nil
}

// GetUserByID IDでユーザーを取得する
func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	model := new(models.UserModel)

	var selector *bun.SelectQuery

	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = r.db.NewSelect()
	}

	selector.Model(model).Where("user_id = ?", id)
	if err := selector.Scan(ctx); err != nil {
		return nil, nil // ユーザーが見つからない場合はnilを返す
	}

	return mapUserModelToDomain(model), nil
}

// CreateUser ユーザーを作成する
func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *user.User) error {
	modelUser := mapUserDomainToModel(user)

	var inserter *bun.InsertQuery
	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		inserter = tx.NewInsert()
	} else {
		inserter = r.db.NewInsert()
	}

	if _, err := inserter.Model(modelUser).Exec(ctx); err != nil {
		return err
	}

	return nil
}

// UpdateUser ユーザーを更新する
func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user *user.User) error {
	modelUser := mapUserDomainToModel(user)

	var updater *bun.UpdateQuery
	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		updater = tx.NewUpdate()
	} else {
		updater = r.db.NewUpdate()
	}

	if _, err := updater.Model(modelUser).WherePK().Exec(ctx); err != nil {
		return err
	}
	return nil
}

// DeleteUser ユーザーを削除する
func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	var deleter *bun.DeleteQuery
	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		deleter = tx.NewDelete()
	} else {
		deleter = r.db.NewDelete()
	}

	if _, err := deleter.Model((*models.UserModel)(nil)).Where("user_id = ?", id).Exec(ctx); err != nil {
		return err
	}
	return nil
}

// mapUserModelToDomain インフラモデルからドメイン型へのマッピング
func mapUserModelToDomain(m *models.UserModel) *user.User {
	return &user.User{
		UserID:       user.UserID(m.UserID),
		LoginID:      user.LoginID(m.LoginID),
		PasswordHash: m.PasswordHash,
		PasswordSalt: m.PasswordSalt,
		Name:         m.Name,
		IsAdmin:      m.IsAdmin,
		CreatedAt:    m.CreatedAt,
	}
}

// mapUserDomainToModel ドメイン型からインフラモデルへのマッピング
func mapUserDomainToModel(u *user.User) *models.UserModel {
	return &models.UserModel{
		UserID:       string(u.UserID),
		LoginID:      string(u.LoginID),
		PasswordHash: u.PasswordHash,
		PasswordSalt: u.PasswordSalt,
		Name:         u.Name,
		IsAdmin:      u.IsAdmin,
		CreatedAt:    u.CreatedAt,
	}
}
