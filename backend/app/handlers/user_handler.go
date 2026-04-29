package handlers

import (
	"net/http"

	"inoUwU/pinu/app/usecases/user"
	"inoUwU/pinu/app/usecases/user/input"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/do"
)

// UserController ユーザーコントローラー
type UserHandler struct {
	userUsecase user.UserService
}

// NewUserHandler ユーザーコントローラーを生成する
func NewUserHandler(i *do.Injector) (*UserHandler, error) {
	userUsecase := do.MustInvoke[user.UserService](i)
	return &UserHandler{
		userUsecase: userUsecase,
	}, nil
}

// GetUserByID ユーザーIDでユーザーを取得します
func (h *UserHandler) GetUserByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "missing user ID"})
	}

	users, err := h.userUsecase.GetAllUsers(c.Context(), &input.GetUsersInput{
		UserId: id,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "cannot retrieve user"})
	}

	return c.Status(http.StatusOK).JSON(users)
}

// GetUsers ユーザー一覧を取得します
func (h *UserHandler) GetUsers(c fiber.Ctx) error {
	i := &input.GetUsersInput{}
	ctx := c.Context()
	result, err := h.userUsecase.GetAllUsers(ctx, i)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve users",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// Register ユーザーを登録します
func (h *UserHandler) Register(c fiber.Ctx) error {
	request := new(input.CreateUserInput)
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	newUser, err := h.userUsecase.CreateUser(c.RequestCtx(), request)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create user"})
	}

	return c.Status(http.StatusOK).JSON(newUser)
}

// Update ユーザー情報を更新し更新後のユーザー情報を返します。
func (h *UserHandler) Update(c fiber.Ctx) error {

	request := new(input.UpdateUserInput)
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	updatedUser, err := h.userUsecase.UpdateUser(c.RequestCtx(), request)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update user"})
	}

	return c.Status(http.StatusOK).JSON(updatedUser)
}

// Delete ユーザーを削除します
func (h *UserHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "missing user ID"})
	}

	_, err := h.userUsecase.DeleteUser(c.RequestCtx(), &input.DeleteUserInput{
		UserId: id,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "cannot delete user"})
	}

	return c.Status(http.StatusNoContent).JSON(nil)
}
