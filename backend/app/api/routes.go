package api

import (
	"fmt"
	"inoUwU/pinu/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

// SetupRoutes APIルートを設定する
func SetupRoutes(app *fiber.App, injector *do.Injector) {
	// APIグループを作成
	api := app.Group("/api")

	// ユーザー関連のルート
	userController, err := controllers.NewUserController(injector)
	if err != nil {
		panic("Failed to create UserController: " + err.Error())
	}

	fmt.Println("Setting up user routes")
	userController.Route(api)
}
