package api

import (
	"github.com/gofiber/swagger"
	"inoUwU/pinu/app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

// SetupRoutes APIルートを設定する
func SetupRoutes(app *fiber.App, injector *do.Injector) {

	// ヘルスチェック
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Pinu API is running!")
	})

	// Swagger UIの設定
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// APIグループを作成
	api := app.Group("/api")

	// ユーザー関連のルート
	userHandler, err := handlers.NewUserHandler(injector)
	if err != nil {
		panic("Failed to create UserHandler: " + err.Error())
	}
	userHandler.Route(api)

	// カテゴリー関連のルート
	categoryHandler, err := handlers.NewCategoryHandler(injector)
	if err != nil {
		panic("Failed to create CategoryHandler: " + err.Error())
	}
	categoryHandler.Route(api)

	// SSEハンドラーの設定
	sseHandler, err := handlers.NewSSEHandler(injector)
	if err != nil {
		panic("Failed to create SSEHandler: " + err.Error())
	}
	sseHandler.Route(api)
}
