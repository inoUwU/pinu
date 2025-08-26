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

	// ユーザー関連のルート
	user := api.Group("/user")
	user.Get("/", userHandler.GetUsers)
	user.Get("/:id", userHandler.GetUserByID)
	user.Post("/register", userHandler.Register)
	user.Post("/update", userHandler.Update)
	user.Delete("/delete", userHandler.Delete)

	// カテゴリー関連のルート
	categoryHandler, err := handlers.NewCategoryHandler(injector)
	if err != nil {
		panic("Failed to create CategoryHandler: " + err.Error())
	}
	categoryGroup := api.Group("/categories")
	{
		categoryGroup.Post("/", categoryHandler.CreateCategory)
		categoryGroup.Get("/", categoryHandler.GetAllCategories)
		categoryGroup.Get("/:id", categoryHandler.GetCategory)
		categoryGroup.Put("/:id", categoryHandler.UpdateCategory)
		categoryGroup.Delete("/:id", categoryHandler.DeleteCategory)
	}

	// メニュー関連のルート
	menuHandler, err := handlers.NewMenuHandler(injector)
	if err != nil {
		panic("Failed to create MenuHandler: " + err.Error())
	}
	menuGroup := api.Group("/menus")
	{
		menuGroup.Post("/", menuHandler.CreateMenu)
		menuGroup.Get("/", menuHandler.GetAllMenus)
		menuGroup.Get("/:id", menuHandler.GetMenu)
		menuGroup.Get("/category/:categoryId", menuHandler.GetMenusByCategory)
		menuGroup.Put("/:id", menuHandler.UpdateMenu)
		menuGroup.Delete("/:id", menuHandler.DeleteMenu)
	}

	// メニューオプション関連のルート
	menuOptionHandler, err := handlers.NewMenuOptionHandler(injector)
	if err != nil {
		panic("Failed to create MenuOptionHandler: " + err.Error())
	}
	menuOptionGroup := api.Group("/menu-options")
	{
		menuOptionGroup.Post("/", menuOptionHandler.CreateMenuOption)
		menuOptionGroup.Get("/", menuOptionHandler.GetAllMenuOptions)
		menuOptionGroup.Get("/:id", menuOptionHandler.GetMenuOption)
		menuOptionGroup.Put("/:id", menuOptionHandler.UpdateMenuOption)
		menuOptionGroup.Delete("/:id", menuOptionHandler.DeleteMenuOption)
	}

	// テーブル関連のルート
	tableHandler, err := handlers.NewTableHandler(injector)
	if err != nil {
		panic("Failed to create TableHandler: " + err.Error())
	}
	tableGroup := api.Group("/tables")
	{
		tableGroup.Post("/", tableHandler.CreateTable)
		tableGroup.Get("/", tableHandler.GetAllTables)
		tableGroup.Get("/:id", tableHandler.GetTable)
		tableGroup.Get("/status/:status", tableHandler.GetTablesByStatus)
		tableGroup.Put("/:id/status", tableHandler.UpdateTableStatus)
		tableGroup.Delete("/:id", tableHandler.DeleteTable)
	}

	// 設定関連のルート
	settingsHandler, err := handlers.NewSettingsHandler(injector)
	if err != nil {
		panic("Failed to create SettingsHandler: " + err.Error())
	}
	settingsGroup := api.Group("/settings")
	{
		settingsGroup.Get("/", settingsHandler.GetAllSettings)
		settingsGroup.Get("/:key", settingsHandler.GetSetting)
		settingsGroup.Post("/", settingsHandler.SetSetting)
		settingsGroup.Delete("/:key", settingsHandler.DeleteSetting)
	}

	// SSEハンドラーの設定
	sseHandler, err := handlers.NewSSEHandler(injector)
	if err != nil {
		panic("Failed to create SSEHandler: " + err.Error())
	}
	sseHandler.Route(api)
}
