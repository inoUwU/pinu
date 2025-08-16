package middleware

import (
	"github.com/samber/do"

	"github.com/uptrace/bun"
	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/handlers"
	"inoUwU/pinu/app/infrastructure/repositories"
	"inoUwU/pinu/app/services"
	"inoUwU/pinu/app/usecases"
)

// Injection 依存性注入コンテナの初期化
func Injection(db *bun.DB, logger port.Logger) (i *do.Injector) {
	injector := do.New()

	// DIコンテナにリポジトリを登録
	do.ProvideNamed(injector, "db", func(i *do.Injector) (*bun.DB, error) {
		return db, nil
	})

	// ロガーをDIコンテナに登録
	do.ProvideNamed(injector, "logger", func(i *do.Injector) (port.Logger, error) {
		return logger, nil
	})

	do.Provide(injector, services.NewSSEService)
	do.Provide(injector, repositories.NewUserRepository)
	do.Provide(injector, usecases.NewUserUsecase)
	do.Provide(injector, handlers.NewUserHandler)
	do.Provide(injector, handlers.NewAuthHandler)
	do.Provide(injector, handlers.NewCategoryHandler)
	return injector
}
