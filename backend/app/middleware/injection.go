package middleware

import (
	"github.com/samber/do"

	"github.com/uptrace/bun"
	"inoUwU/pinu/app/domain/category"
	"inoUwU/pinu/app/domain/menu"
	"inoUwU/pinu/app/domain/menu_option"
	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/domain/settings"
	"inoUwU/pinu/app/domain/table"
	"inoUwU/pinu/app/handlers"
	"inoUwU/pinu/app/infrastructure/repositories"
	categoryRepo "inoUwU/pinu/app/infrastructure/repositories/category"
	menuRepo "inoUwU/pinu/app/infrastructure/repositories/menu"
	menuOptionRepo "inoUwU/pinu/app/infrastructure/repositories/menu_option"
	settingsRepo "inoUwU/pinu/app/infrastructure/repositories/settings"
	tableRepo "inoUwU/pinu/app/infrastructure/repositories/table"
	"inoUwU/pinu/app/services"
	categoryUsecase "inoUwU/pinu/app/usecases/category"
	menuUsecase "inoUwU/pinu/app/usecases/menu"
	menuOptionUsecase "inoUwU/pinu/app/usecases/menu_option"
	settingsUsecase "inoUwU/pinu/app/usecases/settings"
	tableUsecase "inoUwU/pinu/app/usecases/table"
	usecases "inoUwU/pinu/app/usecases/user"
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

	// Services
	do.Provide(injector, services.NewSSEService)

	// User関連
	do.Provide(injector, repositories.NewUserRepository)
	do.Provide(injector, usecases.NewUserUsecase)
	do.Provide(injector, handlers.NewUserHandler)
	do.Provide(injector, handlers.NewAuthHandler)

	// Category関連
	do.Provide(injector, func(i *do.Injector) (category.CategoryRepository, error) {
		db := do.MustInvokeNamed[*bun.DB](i, "db")
		return categoryRepo.NewCategoryRepository(db), nil
	})
	do.Provide(injector, categoryUsecase.NewCategoryUsecase)
	do.Provide(injector, handlers.NewCategoryHandler)

	// Menu関連
	do.Provide(injector, func(i *do.Injector) (menu.MenuRepository, error) {
		db := do.MustInvokeNamed[*bun.DB](i, "db")
		return menuRepo.NewMenuRepository(db), nil
	})
	do.Provide(injector, menuUsecase.NewMenuUsecase)
	do.Provide(injector, handlers.NewMenuHandler)

	// MenuOption関連
	do.Provide(injector, func(i *do.Injector) (menu_option.MenuOptionRepository, error) {
		db := do.MustInvokeNamed[*bun.DB](i, "db")
		return menuOptionRepo.NewMenuOptionRepository(db), nil
	})
	do.Provide(injector, menuOptionUsecase.NewMenuOptionUsecase)
	do.Provide(injector, handlers.NewMenuOptionHandler)

	// Table関連
	do.Provide(injector, func(i *do.Injector) (table.TableRepository, error) {
		return tableRepo.NewTableRepository(i), nil
	})
	do.Provide(injector, tableUsecase.NewTableUsecase)
	do.Provide(injector, handlers.NewTableHandler)

	// Settings関連
	do.Provide(injector, func(i *do.Injector) (settings.SettingsRepository, error) {
		db := do.MustInvokeNamed[*bun.DB](i, "db")
		return settingsRepo.NewSettingsRepository(db), nil
	})
	do.Provide(injector, settingsUsecase.NewSettingsUsecase)
	do.Provide(injector, handlers.NewSettingsHandler)

	return injector
}
