package services

import (
	"github.com/samber/do"
	"inoUwU/pinu/app/usecases"
	"inoUwU/pinu/app/usecases/input"
	"inoUwU/pinu/app/usecases/output"
)

type IUserService interface {
	GetAllUsers() (*output.GetUsersOutput, error)
}

// userService ユーザーサービス
type userService struct {
	userUsecase usecases.IUserUsecase
}

// NewUserService ユーザーサービスを生成する
func NewUserService(i *do.Injector) (IUserService, error) {
	return &userService{
		userUsecase: do.MustInvoke[usecases.IUserUsecase](i),
	}, nil
}

// GetAllUsers 全てのユーザーを取得する
func (us *userService) GetAllUsers() (*output.GetUsersOutput, error) {
	input := &input.GetUsersInput{}
	return us.userUsecase.GetAllUsers(input)
}
