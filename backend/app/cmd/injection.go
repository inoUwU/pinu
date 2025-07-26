package main

import (
	"github.com/samber/do"

	"database/sql"
	"inoUwU/pinu/app/controllers"
	"inoUwU/pinu/app/domain/services"
	"inoUwU/pinu/app/infrastructure/repositories"
	"inoUwU/pinu/app/usecases"
)

// Injection 依存性注入コンテナの初期化
func Injection(db *sql.DB) (i *do.Injector) {
	injector := do.New()

	// DIコンテナにリポジトリを登録
	do.ProvideNamed(injector, "user", func(i *do.Injector) (*sql.DB, error) {
		return db, nil
	})

	do.Provide(injector, repositories.NewUserRepository)
	do.Provide(injector, services.NewUserService)
	do.Provide(injector, usecases.NewUserUsecase)
	do.Provide(injector, controllers.NewUserController)
	return injector
}
