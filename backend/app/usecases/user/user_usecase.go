package usecases

import (
	"context"
	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/domain/user"
	"inoUwU/pinu/app/usecases/user/input"
	"inoUwU/pinu/app/usecases/user/output"
	"inoUwU/pinu/pkg"
	"inoUwU/pinu/pkg/security"

	"github.com/samber/do"
)

// IUserUsecase ユーザーユースケースのインターフェース
type IUserUsecase interface {
	GetAllUsers(ctx context.Context, input *input.GetUsersInput) (*output.GetUsersOutput, error)
	CreateUser(ctx context.Context, input *input.CreateUserInput) (*output.CreateUserOutput, error)
	UpdateUser(ctx context.Context, input *input.UpdateUserInput) (*output.UpdateUserOutput, error)
	DeleteUser(ctx context.Context, input *input.DeleteUserInput) (*output.DeleteUserOutput, error)
}

// UserUsecaseImpl ユーザーユースケースの実装
type UserUsecaseImpl struct {
	userRepo user.IUserRepository
	logger   port.Logger
}

// NewUserUsecase ユーザーユースケースを生成する
func NewUserUsecase(i *do.Injector) (IUserUsecase, error) {
	repository := do.MustInvoke[user.IUserRepository](i)
	logger := do.MustInvokeNamed[port.Logger](i, "logger")

	return &UserUsecaseImpl{
		userRepo: repository,
		logger:   logger,
	}, nil
}

// GetAllUsers 全てのユーザーを取得する
func (u *UserUsecaseImpl) GetAllUsers(ctx context.Context, input *input.GetUsersInput) (*output.GetUsersOutput, error) {
	users, err := u.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return &output.GetUsersOutput{
		Users: users,
		Count: len(users),
	}, nil
}

func (u *UserUsecaseImpl) CreateUser(ctx context.Context, input *input.CreateUserInput) (*output.CreateUserOutput, error) {
	u.logger.Info("creating user", "userID", input.UserId, "", input.Name)
	// TODO: implement transaction

	id, err := pkg.GenerateUUIDv7()
	if err != nil {
		return nil, err
	}

	// Passwordをハッシュ化する
	salt, err := security.GenerateSalt(16)
	if err != nil {
		return nil, err
	}

	hash, err := security.HashPassword("user_input_password", salt)
	if err != nil {
		return nil, err
	}

	modelUser := &user.User{
		USER_ID:       user.UserID(id),
		LOGIN_ID:      user.LoginID(input.LoginId),
		PASSWORD_HASH: hash,
		PASSWORD_SALT: salt,
		NAME:          input.Name,
		IS_ADMIN:      input.IsAdmin,
	}

	if err := u.userRepo.CreateUser(ctx, modelUser); err != nil {
		return nil, err
	}

	// Get Created User
	createdUser, err := u.userRepo.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	u.logger.Info("user created successfully", "userID", id)

	// ユーザー作成のロジックを実装
	return &output.CreateUserOutput{
		User:      *createdUser,
		LoginId:   user.LoginID(createdUser.LOGIN_ID),
		PassWord:  createdUser.PASSWORD_HASH,
		UserId:    user.UserID(createdUser.USER_ID),
		Name:      createdUser.NAME,
		IsAdmin:   createdUser.IS_ADMIN,
		CreatedAt: createdUser.CREATED_AT,
	}, nil
}
func (u *UserUsecaseImpl) UpdateUser(ctx context.Context, input *input.UpdateUserInput) (*output.UpdateUserOutput, error) {
	// ユーザー更新のロジックを実装
	return &output.UpdateUserOutput{}, nil
}
func (u *UserUsecaseImpl) DeleteUser(ctx context.Context, input *input.DeleteUserInput) (*output.DeleteUserOutput, error) {

	// TODO:

	// ユーザー削除のロジックを実装
	return &output.DeleteUserOutput{}, nil
}
