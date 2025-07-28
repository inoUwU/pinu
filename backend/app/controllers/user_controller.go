package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
	"inoUwU/pinu/app/usecases"
	"inoUwU/pinu/app/usecases/input"
)

type IUserController interface {
	GetUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	Route(router fiber.Router) error
}

// UserController ユーザーコントローラー
type UserController struct {
	userUsecase usecases.IUserUsecase
}

// NewUserController ユーザーコントローラーを生成する
func NewUserController(i *do.Injector) (IUserController, error) {
	userUsecase := do.MustInvoke[usecases.IUserUsecase](i)
	return &UserController{
		userUsecase: userUsecase,
	}, nil
}

func (uc *UserController) Route(router fiber.Router) error {
	fmt.Println("Registering user routes")
	router.Get("/users", uc.GetUsers)
	router.Get("/test", uc.GetUserByID)
	return nil
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
func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	input := &input.GetUsersInput{}
	result, err := uc.userUsecase.GetAllUsers(input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve users",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (uc *UserController) GetUserByID(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User retrieved successfully",
	})
}
