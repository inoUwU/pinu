package handlers

import (
	"net/http"

	"inoUwU/pinu/app/usecases/user"
	"inoUwU/pinu/app/usecases/user/input"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type IUserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

// UserController ユーザーコントローラー
type UserHandler struct {
	userUsecase usecases.IUserUsecase
}

// NewUserHandler ユーザーコントローラーを生成する
func NewUserHandler(i *do.Injector) (IUserHandler, error) {
	userUsecase := do.MustInvoke[usecases.IUserUsecase](i)
	return &UserHandler{
		userUsecase: userUsecase,
	}, nil
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User retrieved successfully",
	})
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	request := new(input.CreateUserInput)
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	output, err := h.userUsecase.CreateUser(c.UserContext(), request)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create user"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result":  output,
		"message": "User registered successfully",
	})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

// GetUsers ユーザー一覧を取得するAPIハンドラー
// @Summary ユーザー一覧取得
// @Description 全てのユーザーの一覧を取得する
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} output.GetUsersOutput
// @Failure 500 {object} map[string]string
// @Router /api/users [get]
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	input := &input.GetUsersInput{}
	ctx := c.UserContext()
	result, err := h.userUsecase.GetAllUsers(ctx, input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve users",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}
