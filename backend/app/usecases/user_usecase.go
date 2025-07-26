package usecases

import (
	"github.com/samber/do"
	"inoUwU/pinu/app/domain/repositories"
	"inoUwU/pinu/app/usecases/input"
	"inoUwU/pinu/app/usecases/output"
)

// IUserUsecase ユーザーユースケースのインターフェース
type IUserUsecase interface {
	GetAllUsers(input *input.GetUsersInput) (*output.GetUsersOutput, error)
}

// UserUsecaseImpl ユーザーユースケースの実装
type UserUsecaseImpl struct {
	userRepo repositories.IUserRepository
}

// NewUserUsecase ユーザーユースケースを生成する
func NewUserUsecase(i *do.Injector) (*UserUsecaseImpl, error) {
	repository := do.MustInvoke[repositories.IUserRepository](i)
	return &UserUsecaseImpl{
		userRepo: repository,
	}, nil
}

// GetAllUsers 全てのユーザーを取得する
func (u *UserUsecaseImpl) GetAllUsers(input *input.GetUsersInput) (*output.GetUsersOutput, error) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return &output.GetUsersOutput{
		Users: users,
		Count: len(users),
	}, nil
}
