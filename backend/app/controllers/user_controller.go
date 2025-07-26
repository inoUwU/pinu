package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
	"inoUwU/pinu/app/domain/services"
)

// UserController ユーザーコントローラー
type UserController struct {
	userService services.UserService
}

// NewUserController ユーザーコントローラーを生成する
func NewUserController(i *do.Injector) (*UserController, error) {
	userService := do.MustInvoke[services.UserService](i)
	return &UserController{
		userService: userService,
	}, nil
}

func (uc *UserController) Route(router fiber.Router) {
	router.Get("/users", uc.GetUsers)
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
	result, err := uc.userService.GetAllUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve users",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}
