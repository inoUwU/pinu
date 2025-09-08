package repositories

import (
	"context"
	"inoUwU/pinu/app/domain/user"
	"inoUwU/pinu/app/infrastructure/models"

	"inoUwU/pinu/app/infrastructure/ctx"

	"github.com/samber/do"
	"github.com/uptrace/bun"
)

// ユーザーリポジトリの実装（アダプター）
type UserRepositoryImpl struct {
	db *bun.DB
}

// ユーザーリポジトリの実装を生成する
func NewUserRepository(i *do.Injector) (user.IUserRepository, error) {
	db := do.MustInvokeNamed[*bun.DB](i, "db")
	return &UserRepositoryImpl{db: db}, nil
}

// GetAllUsers 全てのユーザーを取得する
func (r *UserRepositoryImpl) GetAllUsers(ctx context.Context) ([]user.User, error) {
	tmp_users := make([]models.User, 0)
	if err := r.db.NewSelect().Model(&tmp_users).Scan(ctx); err != nil {
		return nil, err
	}

	users := make([]user.User, len(tmp_users))
	for i := 0; i < len(tmp_users); i++ {
		model := tmp_users[i]
		users[i] = user.User{
			USER_ID:       user.UserID(model.USER_ID),
			LOGIN_ID:      user.LoginID(model.LOGIN_ID),
			PASSWORD_HASH: model.PASSWORD_HASH,
			NAME:          model.NAME,
			IS_ADMIN:      model.IS_ADMIN,
			CREATED_AT:    model.CREATED_AT,
		}
	}

	return users, nil
}

// ログインIDでユーザーを取得する
func (r *UserRepositoryImpl) GetUserByLoginID(ctx context.Context, loginID string) (*user.User, error) {
	user := new(user.User)

	var selecter *bun.SelectQuery

	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		// トランザクションオブジェクトが存在する場合はそれを使用
		selecter = tx.NewSelect()
	} else {
		// トランザクションオブジェクトが存在しない場合はDBオブジェクトを使用
		selecter = r.db.NewSelect()
	}

	// 実行処理は共通化
	selecter.Model(user).Where("login_id = ?", loginID)
	if err := selecter.Scan(ctx); err != nil {
		return nil, err
	}
	return user, nil
}

// IDでユーザーを取得する
func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	user := new(user.User)

	var selecter *bun.SelectQuery

	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		// トランザクションオブジェクトが存在する場合はそれを使用
		selecter = tx.NewSelect()
	} else {
		// トランザクションオブジェクトが存在しない場合はDBオブジェクトを使用
		selecter = r.db.NewSelect()
	}

	// 実行処理は共通化
	selecter.Model(user).Where("id = ?", id)
	if err := selecter.Scan(ctx); err != nil {
		return nil, nil // ユーザーが見つからない場合はnilを返す
	}
	return user, nil
}

// ユーザーを作成する
func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *user.User) error {
	modelUser := &models.User{
		USER_ID:       string(user.USER_ID),
		LOGIN_ID:      string(user.LOGIN_ID),
		PASSWORD_HASH: user.PASSWORD_HASH,
		PASSWORD_SALT: user.PASSWORD_SALT,
		NAME:          user.NAME,
		IS_ADMIN:      user.IS_ADMIN,
		CREATED_AT:    user.CREATED_AT,
	}

	var inserter *bun.InsertQuery
	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		// トランザクションオブジェクトが存在する場合はそれを使用
		inserter = tx.NewInsert()
	} else {
		// トランザクションオブジェクトが存在しない場合はDBオブジェクトを使用
		inserter = r.db.NewInsert()
	}

	if _, err := inserter.Model(modelUser).Exec(ctx); err != nil {
		return err
	}

	return nil
}

// ユーザーを更新する
func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user *user.User) error {
	modelUser := &models.User{
		USER_ID:       string(user.USER_ID),
		LOGIN_ID:      string(user.LOGIN_ID),
		PASSWORD_HASH: user.PASSWORD_HASH,
		PASSWORD_SALT: user.PASSWORD_SALT,
		NAME:          user.NAME,
		IS_ADMIN:      user.IS_ADMIN,
		CREATED_AT:    user.CREATED_AT,
	}

	var inserter *bun.UpdateQuery
	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		// トランザクションオブジェクトが存在する場合はそれを使用
		inserter = tx.NewUpdate()
	} else {
		// トランザクションオブジェクトが存在しない場合はDBオブジェクトを使用
		inserter = r.db.NewUpdate()
	}

	if _, err := inserter.Model(modelUser).WherePK().Exec(ctx); err != nil {
		return err
	}
	return nil
}

// ユーザーを削除する
func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	var deleter *bun.DeleteQuery
	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		// トランザクションオブジェクトが存在する場合はそれを使用
		deleter = tx.NewDelete()
	} else {
		// トランザクションオブジェクトが存在しない場合はDBオブジェクトを使用
		deleter = r.db.NewDelete()
	}

	if _, err := deleter.Where("user_id = ?", id).Exec(ctx); err != nil {
		return err
	}
	return nil
}
