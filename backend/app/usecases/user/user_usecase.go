package user

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

// UserService ユーザーユースケースのポート
type UserService interface {
	GetAllUsers(ctx context.Context, input *input.GetUsersInput) (*output.GetUsersOutput, error)
	CreateUser(ctx context.Context, input *input.CreateUserInput) (*output.CreateUserOutput, error)
	UpdateUser(ctx context.Context, input *input.UpdateUserInput) (*output.UpdateUserOutput, error)
	DeleteUser(ctx context.Context, input *input.DeleteUserInput) (*output.DeleteUserOutput, error)
}

// UserUsecaseImpl ユーザーユースケースの実装
type UserUsecaseImpl struct {
	unitOfWork port.UnitOfWork
	userRepo   user.UserRepository
	logger     port.Logger
}

// NewUserUsecase ユーザーユースケースを生成する
func NewUserUsecase(i *do.Injector) (UserService, error) {
	repository := do.MustInvoke[user.UserRepository](i)
	logger := do.MustInvokeNamed[port.Logger](i, "logger")
	unitOfWork := do.MustInvokeNamed[port.UnitOfWork](i, "uow")

	return &UserUsecaseImpl{
		userRepo:   repository,
		logger:     logger,
		unitOfWork: unitOfWork,
	}, nil
}

// GetAllUsers 全てのユーザーを取得する
func (u *UserUsecaseImpl) GetAllUsers(ctx context.Context, input *input.GetUsersInput) (*output.GetUsersOutput, error) {
	var users []user.User = nil

	// ユーザー一覧を取得するトランザクション処理
	err := u.unitOfWork.Run(ctx, func(ctx context.Context) error {
		// userRepoを使ってユーザー一覧を取得
		var err error
		users, err = u.userRepo.GetAllUsers(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &output.GetUsersOutput{
		Users: users,
	}, nil
}

// CreateUser 新しいユーザーを作成し登録後のユーザー情報を返す
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

	hash, err := security.HashPassword(input.PassWord, salt)
	if err != nil {
		return nil, err
	}

	modelUser := &user.User{
		UserID:       user.UserID(id),
		LoginID:      user.LoginID(input.LoginId),
		PasswordHash: hash,
		PasswordSalt: salt,
		Name:         input.Name,
		IsAdmin:      input.IsAdmin,
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

	return &output.CreateUserOutput{
		User:      *createdUser,
		LoginId:   user.LoginID(createdUser.LoginID),
		PassWord:  createdUser.PasswordHash,
		UserId:    user.UserID(createdUser.UserID),
		Name:      createdUser.Name,
		IsAdmin:   createdUser.IsAdmin,
		CreatedAt: createdUser.CreatedAt,
	}, nil
}

// UpdateUser ユーザーを更新する
func (u *UserUsecaseImpl) UpdateUser(ctx context.Context, input *input.UpdateUserInput) (*output.UpdateUserOutput, error) {
	// ユーザー更新のロジックを実装
	return &output.UpdateUserOutput{}, nil
}

// DeleteUser ユーザーを削除する
func (u *UserUsecaseImpl) DeleteUser(ctx context.Context, input *input.DeleteUserInput) (*output.DeleteUserOutput, error) {

	// TODO:

	// ユーザー削除のロジックを実装
	return &output.DeleteUserOutput{}, nil
}
