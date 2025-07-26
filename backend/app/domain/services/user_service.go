package services

import (
	"github.com/samber/do"
	"inoUwU/pinu/app/usecases"
	"inoUwU/pinu/app/usecases/input"
	"inoUwU/pinu/app/usecases/output"
)

// UserService ユーザーサービス
type UserService struct {
	userUsecase usecases.IUserUsecase
}

// NewUserService ユーザーサービスを生成する
func NewUserService(i *do.Injector) (*UserService, error) {
	return &UserService{}, nil
}

// GetAllUsers 全てのユーザーを取得する
func (s *UserService) GetAllUsers() (*output.GetUsersOutput, error) {
	input := &input.GetUsersInput{}
	return s.userUsecase.GetAllUsers(input)
}
