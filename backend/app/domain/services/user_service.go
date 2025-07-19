package services

import (
	"inoUwU/pinu/app/usecases"
	"inoUwU/pinu/app/usecases/input"
	"inoUwU/pinu/app/usecases/output"
)

// UserService ユーザーサービス
type UserService struct {
	userUsecase usecases.UserUsecase
}

// NewUserService ユーザーサービスを生成する
func NewUserService(userUsecase usecases.UserUsecase) *UserService {
	return &UserService{
		userUsecase: userUsecase,
	}
}

// GetAllUsers 全てのユーザーを取得する
func (s *UserService) GetAllUsers() (*output.GetUsersOutput, error) {
	input := &input.GetUsersInput{}
	return s.userUsecase.GetAllUsers(input)
}
