package api

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"inoUwU/pinu/app/controllers"
	"inoUwU/pinu/app/domain/services"
	"inoUwU/pinu/app/infrastructure/repositories"
	"inoUwU/pinu/app/usecases"
)

// SetupRoutes APIルートを設定する
func SetupRoutes(app *fiber.App, db *sql.DB) {
	// 依存関係の注入（DI）
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userService := services.NewUserService(userUsecase)
	userController := controllers.NewUserController(userService)

	// APIグループを作成
	api := app.Group("/api")

	// ユーザー関連のルート
	users := api.Group("/users")
	users.Get("/", userController.GetUsers)
}
