package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
	"inoUwU/pinu/app/usecases"
	"inoUwU/pinu/app/usecases/input"
)

type IUserHandler interface {
	Handler
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

func (uh *UserHandler) Route(router fiber.Router) error {
	fmt.Println("Registering user routes")
	user := router.Group("/user")
	user.Get("/", uh.GetUsers)
	user.Get("/:id", uh.GetUserByID)
	user.Post("/register", uh.Register)
	user.Post("/update", uh.Update)
	user.Delete("/delete", uh.Delete)
	return nil
}

func (uh *UserHandler) GetUserByID(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User retrieved successfully",
	})
}

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func (uh *UserHandler) Update(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func (uh *UserHandler) Delete(c *fiber.Ctx) error {
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
func (uh *UserHandler) GetUsers(c *fiber.Ctx) error {
	input := &input.GetUsersInput{}
	result, err := uh.userUsecase.GetAllUsers(input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve users",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}
