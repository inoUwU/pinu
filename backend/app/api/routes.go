package api

import (
	"inoUwU/pinu/app/handlers"

	swagger "github.com/gofiber/contrib/v3/swaggo"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/do"
)

// SetupRoutes APIルートを設定する
func SetupRoutes(app *fiber.App, injector *do.Injector) {

	// ヘルスチェック
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Pinu API is running!")
	})

	// Swagger UIの設定
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// APIグループを作成
	api := app.Group("/api")

	// 認証関連のルート
	authHandler, err := handlers.NewAuthHandler(injector)
	if err != nil {
		panic("Failed to create AuthHandler: " + err.Error())
	}
	authGroup := api.Group("/auth")
	{
		authGroup.Post("/login", authHandler.Login)
		authGroup.Post("/logout", authHandler.Logout)
		authGroup.Post("/renew", authHandler.RenewAccessToken)
		authGroup.Post("/revoke", authHandler.RevokeSession)
	}

	// ユーザー関連のルート
	userHandler, err := handlers.NewUserHandler(injector)
	if err != nil {
		panic("Failed to create UserHandler: " + err.Error())
	}
	userGroup := api.Group("/user")
	{
		userGroup.Get("/", userHandler.GetUsers)
		userGroup.Get("/:id", userHandler.GetUserByID)
		userGroup.Post("/register", userHandler.Register)
		userGroup.Post("/update", userHandler.Update)
		userGroup.Delete("/delete", userHandler.Delete)
	}

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
		tableGroup.Post("/:id/checkout", tableHandler.Checkout)
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

	// 注文関連のルート
	orderHandler, err := handlers.NewOrderHandler(injector)
	if err != nil {
		panic("Failed to create OrderHandler: " + err.Error())
	}
	orderGroup := api.Group("/orders")
	{
		orderGroup.Get("/", orderHandler.GetAllOrders)
		orderGroup.Post("/", orderHandler.Order)
	}

	// 統計関連のルート
	analyticsHandler, err := handlers.NewAnalyticsHandler(injector)
	if err != nil {
		panic("Failed to create AnalyticsHandler: " + err.Error())
	}
	analyticsGroup := api.Group("/analytics")
	{
		analyticsGroup.Get("/", analyticsHandler.GetAnalyticsData)
		analyticsGroup.Get("/kpi", analyticsHandler.GetKPISummary)
		analyticsGroup.Get("/top-menus", analyticsHandler.GetTopMenus)
		analyticsGroup.Get("/category-sales", analyticsHandler.GetCategorySales)
		analyticsGroup.Get("/menu-performance", analyticsHandler.GetMenuPerformance)
		analyticsGroup.Get("/daily-sales", analyticsHandler.GetDailySales)
		analyticsGroup.Get("/menu-daily-trends", analyticsHandler.GetMenuDailyTrends)
	}

	// TODO: SSEハンドラーではなく注文画面と支払い画面に分割する
	// SSEハンドラーの設定
	sseHandler, err := handlers.NewSSEHandler(injector)
	if err != nil {
		panic("Failed to create SSEHandler: " + err.Error())
	}
	if err := sseHandler.Route(api); err != nil {
		panic("Failed to register SSE routes: " + err.Error())
	}
}
