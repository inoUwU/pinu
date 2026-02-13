package auth

import (
	"context"
	"errors"
	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/domain/session"
	"inoUwU/pinu/app/domain/user"
	"inoUwU/pinu/app/usecases/auth/input"
	"inoUwU/pinu/app/usecases/auth/output"
	"inoUwU/pinu/pkg/security"
	"inoUwU/pinu/pkg/security/token"
	"time"

	"github.com/samber/do"
)

type AuthService interface {
	Login(ctx context.Context, input *input.Login) (*output.Login, error)
	Logout(ctx context.Context, token string) error
	RenewAccessToken(ctx context.Context, input *input.RenewAccessToken) (*output.RenewAccessToken, error)
	RevokeSession(ctx context.Context, input input.RevokeSession) error
}

type AuthUsecaseImpl struct {
	unitOfWork port.UnitOfWork
	authRepo   session.SessionStore
	userRepo   user.UserStore
	logger     port.Logger
	tokenMaker *token.JWTMaker
}

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidPassword = errors.New("invalid password")

// NewAuthUsecase 認証ユースケースを生成する
func NewAuthUsecase(i *do.Injector) (AuthService, error) {
	repository := do.MustInvoke[session.SessionStore](i)
	logger := do.MustInvokeNamed[port.Logger](i, "logger")
	userRepo := do.MustInvoke[user.UserStore](i)
	unitOfWork := do.MustInvokeNamed[port.UnitOfWork](i, "uow")
	tokenMaker := do.MustInvokeNamed[*token.JWTMaker](i, "jwtMaker")

	return &AuthUsecaseImpl{
		authRepo:   repository,
		logger:     logger,
		userRepo:   userRepo,
		unitOfWork: unitOfWork,
		tokenMaker: tokenMaker,
	}, nil
}

// Login ログインを実行する
func (u *AuthUsecaseImpl) Login(ctx context.Context, input *input.Login) (*output.Login, error) {
	u.logger.Info("login user", "login", input.LoginId)

	// ユーザーの存在をチェック
	loginUser, err := u.userRepo.GetUserByLoginID(ctx, string(input.LoginId))

	if err != nil {
		u.logger.Error("failed to get user", "error", err)
		return nil, err
	}

	if loginUser == nil {
		u.logger.Error("user not found", "login", input.LoginId)
		return nil, ErrUserNotFound
	}

	isValid := security.VerifyPassword(input.Password, loginUser.PASSWORD_HASH, loginUser.PASSWORD_SALT)

	if !isValid {
		u.logger.Error("invalid password", "login", input.LoginId)
		return nil, ErrInvalidPassword
	}

	accessToken, accessClaims, err := u.tokenMaker.GenerateToken(loginUser, 15*time.Minute) // 15分
	if err != nil {
		u.logger.Error("failed to generate token", "error", err)
		return nil, err
	}

	// リフレッシュトークンの作成
	refreshToken, refreshClaims, err := u.tokenMaker.GenerateToken(loginUser, 7*24*time.Hour) // 7日
	if err != nil {
		u.logger.Error("failed to generate refresh token", "error", err)
		return nil, err
	}

	err = u.unitOfWork.Run(ctx, func(ctx context.Context) error {
		modelSession := &session.Session{
			SessionID:    refreshClaims.RegisteredClaims.ID,
			UserID:       loginUser.USER_ID,
			RefreshToken: refreshToken,
			IsRevoked:    false,
			ExpiresAt:    refreshClaims.RegisteredClaims.ExpiresAt.Time,
		}

		if err := u.authRepo.CreateSession(ctx, modelSession); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	res := output.Login{
		RefreshToken:          refreshToken,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshClaims.ExpiresAt.Time,
		User: output.LoginUserRes{
			UserId:   string(loginUser.USER_ID),
			LoginId:  string(loginUser.LOGIN_ID),
			UserName: loginUser.NAME,
			IsAdmin:  loginUser.IS_ADMIN,
		},
	}

	return &res, nil
}

// Logout ログアウトを実行する
func (u *AuthUsecaseImpl) Logout(ctx context.Context, token string) error {
	err := u.authRepo.DeleteSession(ctx, token)
	if err != nil {
		u.logger.Error("failed to logout", "error", err)
		return err
	}
	return nil
}

// RenewAccessToken リフレッシュトークンを使ってアクセストークンを更新する
func (u *AuthUsecaseImpl) RenewAccessToken(ctx context.Context, input *input.RenewAccessToken) (*output.RenewAccessToken, error) {
	// トークンの検証
	refreshClaims, err := u.tokenMaker.VerifyToken(input.RefreshToken)
	if err != nil {
		u.logger.Error("failed to verify token", "error", err)
		return nil, err
	}

	// セッションの取得
	userSession, err := u.authRepo.GetSessionByID(ctx, refreshClaims.RegisteredClaims.ID)
	if err != nil {
		u.logger.Error("failed to get session", "error", err)
		return nil, err
	}

	if userSession == nil || userSession.IsRevoked {
		u.logger.Error("session not found or revoked", "session_id", refreshClaims.RegisteredClaims.ID)
		return nil, errors.New("session not found or revoked")
	}

	if string(userSession.UserID) != refreshClaims.UserId {
		u.logger.Error("session user id does not match token user id", "session_user_id", userSession.UserID, "token_user_id", refreshClaims.UserId)
		return nil, errors.New("session user id does not match token user id")
	}

	// 新しいアクセストークンの発行
	accessUser, err := u.userRepo.GetUserByID(ctx, refreshClaims.UserId)
	if err != nil {
		u.logger.Error("failed to get user by id", "error", err)
		return nil, err
	}

	newAccessToken, newAccessClaims, err := u.tokenMaker.GenerateToken(accessUser, 15*time.Minute) // 15分
	if err != nil {
		u.logger.Error("failed to generate new access token", "error", err)
		return nil, err
	}

	return &output.RenewAccessToken{
		AccessToken:          newAccessToken,
		AccessTokenExpiresAt: newAccessClaims.ExpiresAt.Time,
	}, nil
}

func (u *AuthUsecaseImpl) RevokeSession(ctx context.Context, input input.RevokeSession) error {
	userSession, err := u.authRepo.GetSessionByID(ctx, input.SessionId)
	if err != nil {
		u.logger.Error("failed to get session by id", "error", err)
		return err
	}

	if userSession == nil {
		u.logger.Error("session not found", "session_id", input.SessionId)
		return errors.New("session not found")
	}

	err = u.authRepo.RevokeSessionByID(ctx, input.SessionId)
	if err != nil {
		u.logger.Error("failed to revoke session by id", "error", err)
		return err
	}
	return nil
}
