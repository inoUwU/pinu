package main

import (
	"inoUwU/pinu/app/domain/analytics"
	"inoUwU/pinu/app/domain/category"
	"inoUwU/pinu/app/domain/menu"
	"inoUwU/pinu/app/domain/menu_option"
	"inoUwU/pinu/app/domain/order"
	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/domain/settings"
	"inoUwU/pinu/app/domain/table"
	"inoUwU/pinu/app/handlers"
	"inoUwU/pinu/app/infrastructure/repositories"
	analyticsRepo "inoUwU/pinu/app/infrastructure/repositories/analytics"
	categoryRepo "inoUwU/pinu/app/infrastructure/repositories/category"
	menuRepo "inoUwU/pinu/app/infrastructure/repositories/menu"
	menuOptionRepo "inoUwU/pinu/app/infrastructure/repositories/menu_option"
	orderRepo "inoUwU/pinu/app/infrastructure/repositories/order"
	"inoUwU/pinu/app/infrastructure/repositories/session"
	settingsRepo "inoUwU/pinu/app/infrastructure/repositories/settings"
	tableRepo "inoUwU/pinu/app/infrastructure/repositories/table"
	userRepo "inoUwU/pinu/app/infrastructure/repositories/user"
	"inoUwU/pinu/app/infrastructure/sse"
	analyticsUsecase "inoUwU/pinu/app/usecases/analytics"
	"inoUwU/pinu/app/usecases/auth"
	categoryUsecase "inoUwU/pinu/app/usecases/category"
	menuUsecase "inoUwU/pinu/app/usecases/menu"
	menuOptionUsecase "inoUwU/pinu/app/usecases/menu_option"
	orderUsecase "inoUwU/pinu/app/usecases/order"
	settingsUsecase "inoUwU/pinu/app/usecases/settings"
	tableUsecase "inoUwU/pinu/app/usecases/table"
	usecases "inoUwU/pinu/app/usecases/user"
	"inoUwU/pinu/pkg/security/token"

	"github.com/samber/do"
	"github.com/uptrace/bun"
)

func buildInjector(db *bun.DB, logger port.Logger, secret string) *do.Injector {
	injector := do.New()

	do.ProvideNamed(injector, "db", func(i *do.Injector) (*bun.DB, error) {
		return db, nil
	})

	do.ProvideNamed(injector, "logger", func(i *do.Injector) (port.Logger, error) {
		return logger, nil
	})

	do.ProvideNamed(injector, "jwtMaker", func(i *do.Injector) (port.TokenMaker, error) {
		jwtMaker, err := token.NewJwtMaker(secret)
		if err != nil {
			return nil, err
		}
		return jwtMaker, nil
	})

	do.ProvideNamed(injector, "uow", func(i *do.Injector) (port.UnitOfWork, error) {
		tx := repositories.NewTxRepository(db)
		return tx, nil
	})

	do.Provide(injector, session.NewSessionRepository)
	do.Provide(injector, auth.NewAuthUsecase)
	do.Provide(injector, handlers.NewAuthHandler)

	do.Provide(injector, sse.NewSSEBroker)

	do.Provide(injector, userRepo.NewUserRepository)
	do.Provide(injector, usecases.NewUserUsecase)
	do.Provide(injector, handlers.NewUserHandler)

	do.Provide(injector, func(i *do.Injector) (category.CategoryRepository, error) {
		db := do.MustInvokeNamed[*bun.DB](i, "db")
		return categoryRepo.NewCategoryRepository(db), nil
	})
	do.Provide(injector, categoryUsecase.NewCategoryUsecase)
	do.Provide(injector, handlers.NewCategoryHandler)

	do.Provide(injector, func(i *do.Injector) (menu.MenuRepository, error) {
		db := do.MustInvokeNamed[*bun.DB](i, "db")
		return menuRepo.NewMenuRepository(db), nil
	})
	do.Provide(injector, menuUsecase.NewMenuUsecase)
	do.Provide(injector, handlers.NewMenuHandler)

	do.Provide(injector, func(i *do.Injector) (menu_option.MenuOptionRepository, error) {
		db := do.MustInvokeNamed[*bun.DB](i, "db")
		return menuOptionRepo.NewMenuOptionRepository(db), nil
	})
	do.Provide(injector, menuOptionUsecase.NewMenuOptionUsecase)
	do.Provide(injector, handlers.NewMenuOptionHandler)

	do.Provide(injector, func(i *do.Injector) (table.TableRepository, error) {
		return tableRepo.NewTableRepository(i), nil
	})
	do.Provide(injector, tableUsecase.NewTableUsecase)
	do.Provide(injector, handlers.NewTableHandler)

	do.Provide(injector, func(i *do.Injector) (settings.SettingsRepository, error) {
		db := do.MustInvokeNamed[*bun.DB](i, "db")
		return settingsRepo.NewSettingsRepository(db), nil
	})
	do.Provide(injector, settingsUsecase.NewSettingsUsecase)
	do.Provide(injector, handlers.NewSettingsHandler)

	do.Provide(injector, func(i *do.Injector) (analytics.AnalyticsRepository, error) {
		return analyticsRepo.NewAnalyticsRepository(i)
	})
	do.Provide(injector, analyticsUsecase.NewAnalyticsUsecase)
	do.Provide(injector, handlers.NewAnalyticsHandler)

	do.Provide(injector, func(i *do.Injector) (order.OrderRepository, error) {
		return orderRepo.NewOrderRepository(i)
	})
	do.Provide(injector, orderUsecase.NewOrderUsecase)
	do.Provide(injector, handlers.NewOrderHandler)

	return injector
}
