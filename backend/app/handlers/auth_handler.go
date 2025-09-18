package handlers

import (
	"inoUwU/pinu/app/usecases/auth"
	"inoUwU/pinu/app/usecases/auth/input"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type IAuthHandler interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	RenewAccessToken(c *fiber.Ctx) error
	RevokeSession(c *fiber.Ctx) error
}

// AuthHandler UserAuthHandler ユーザー認証ハンドラー
type AuthHandler struct {
	authUsecase auth.IAuthUsecase
}

// NewAuthHandler ユーザー認証ハンドラーを生成する
func NewAuthHandler(i *do.Injector) (IAuthHandler, error) {
	authUsecase := do.MustInvoke[auth.IAuthUsecase](i)

	return &AuthHandler{
		authUsecase: authUsecase,
	}, nil
}

// Login ログイン
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	request := new(input.Login)
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	res, err := h.authUsecase.Login(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// Logout  ログアウト
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "missing session ID"})
	}

	err := h.authUsecase.Logout(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to logout", "message": err.Error()})
	}

	return c.Status(http.StatusNoContent).JSON(fiber.Map{"message": "successfully logged out"})
}

// RenewAccessToken アクセストークンの更新
func (h *AuthHandler) RenewAccessToken(c *fiber.Ctx) error {
	request := new(input.RenewAccessToken)
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	res, err := h.authUsecase.RenewAccessToken(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// RevokeSession セッションの無効化
func (h *AuthHandler) RevokeSession(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "missing session ID"})
	}

	// user_idはミドルウェアでセットされていることを想定
	if c.Locals("user_id") == nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	err := h.authUsecase.RevokeSession(c.Context(), input.RevokeSession{
		SessionId: id,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to revoke session", "message": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "successfully revoked session"})
}
