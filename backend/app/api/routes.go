package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
	"inoUwU/pinu/app/controllers"
)

// SetupRoutes APIルートを設定する
func SetupRoutes(app *fiber.App, injector *do.Injector) {
	// APIグループを作成
	api := app.Group("/api")

	// ユーザー関連のルート
	users := api.Group("/users")
	userController, err := controllers.NewUserController(injector)
	if err != nil {
		panic("Failed to create UserController: " + err.Error())
	}

	userController.Route(users)
}
