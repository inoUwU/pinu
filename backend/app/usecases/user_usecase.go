package usecases

import (
	"inoUwU/pinu/app/domain/repositories"
	"inoUwU/pinu/app/usecases/input"
	"inoUwU/pinu/app/usecases/output"
)

// UserUsecase ユーザーユースケースのインターフェース
type UserUsecase interface {
	GetAllUsers(input *input.GetUsersInput) (*output.GetUsersOutput, error)
}

// UserUsecaseImpl ユーザーユースケースの実装
type UserUsecaseImpl struct {
	userRepo repositories.UserRepository
}

// NewUserUsecase ユーザーユースケースを生成する
func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &UserUsecaseImpl{
		userRepo: userRepo,
	}
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
