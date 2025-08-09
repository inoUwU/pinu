package usecases

import (
	"context"
	"github.com/samber/do"
	"inoUwU/pinu/app/domain/repositories"
	"inoUwU/pinu/app/usecases/input"
	"inoUwU/pinu/app/usecases/output"
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
	userRepo repositories.IUserRepository
}

// NewUserUsecase ユーザーユースケースを生成する
func NewUserUsecase(i *do.Injector) (IUserUsecase, error) {
	repository := do.MustInvoke[repositories.IUserRepository](i)
	return &UserUsecaseImpl{
		userRepo: repository,
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
	// ユーザー作成のロジックを実装
	return &output.CreateUserOutput{}, nil
}
func (u *UserUsecaseImpl) UpdateUser(ctx context.Context, input *input.UpdateUserInput) (*output.UpdateUserOutput, error) {
	// ユーザー更新のロジックを実装
	return &output.UpdateUserOutput{}, nil
}
func (u *UserUsecaseImpl) DeleteUser(ctx context.Context, input *input.DeleteUserInput) (*output.DeleteUserOutput, error) {
	// ユーザー削除のロジックを実装
	return &output.DeleteUserOutput{}, nil
}
