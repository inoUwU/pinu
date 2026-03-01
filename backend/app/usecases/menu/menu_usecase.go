package menu

import (
	"context"
	"time"

	"github.com/samber/do"

	"inoUwU/pinu/app/domain/menu"
	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/usecases/menu/input"
	"inoUwU/pinu/app/usecases/menu/output"
)

// MenuService メニューユースケースのポート
type MenuService interface {
	GetAllMenus(ctx context.Context, input *input.GetMenusInput) (*output.GetMenusOutput, error)
	GetMenuByID(ctx context.Context, input *input.GetMenuByIDInput) (*output.GetMenuByIDOutput, error)
	GetMenusByCategory(ctx context.Context, input *input.GetMenusByCategoryInput) (*output.GetMenusByCategoryOutput, error)
	CreateMenu(ctx context.Context, input *input.CreateMenuInput) (*output.CreateMenuOutput, error)
	UpdateMenu(ctx context.Context, input *input.UpdateMenuInput) (*output.UpdateMenuOutput, error)
	DeleteMenu(ctx context.Context, input *input.DeleteMenuInput) (*output.DeleteMenuOutput, error)
}

// MenuUsecaseImpl メニューユースケースの実装
type MenuUsecaseImpl struct {
	menuRepo menu.MenuRepository
	logger   port.Logger
}

// NewMenuUsecase メニューユースケースを生成する
func NewMenuUsecase(i *do.Injector) (MenuService, error) {
	repository := do.MustInvoke[menu.MenuRepository](i)
	logger := do.MustInvokeNamed[port.Logger](i, "logger")

	return &MenuUsecaseImpl{
		menuRepo: repository,
		logger:   logger,
	}, nil
}

// GetAllMenus 全てのメニューを取得する
func (u *MenuUsecaseImpl) GetAllMenus(ctx context.Context, input *input.GetMenusInput) (*output.GetMenusOutput, error) {
	u.logger.Info("getting all menus")

	menus, err := u.menuRepo.GetAll(ctx)
	if err != nil {
		u.logger.Error("failed to get all menus", "error", err)
		return nil, err
	}

	u.logger.Info("successfully retrieved menus", "count", len(menus))

	return &output.GetMenusOutput{
		Menus: menus,
		Count: len(menus),
	}, nil
}

// GetMenuByID 指定IDのメニューを取得する
func (u *MenuUsecaseImpl) GetMenuByID(ctx context.Context, input *input.GetMenuByIDInput) (*output.GetMenuByIDOutput, error) {
	u.logger.Info("getting menu by ID", "menuID", input.MenuID)

	menu, err := u.menuRepo.GetByID(ctx, input.MenuID)
	if err != nil {
		u.logger.Error("failed to get menu by ID", "menuID", input.MenuID, "error", err)
		return nil, err
	}

	if menu == nil {
		u.logger.Warn("menu not found", "menuID", input.MenuID)
		return &output.GetMenuByIDOutput{
			Menu: nil,
		}, nil
	}

	u.logger.Info("successfully retrieved menu", "menuID", input.MenuID)

	return &output.GetMenuByIDOutput{
		Menu: menu,
	}, nil
}

// GetMenusByCategory 指定カテゴリのメニューを取得する
func (u *MenuUsecaseImpl) GetMenusByCategory(ctx context.Context, input *input.GetMenusByCategoryInput) (*output.GetMenusByCategoryOutput, error) {
	u.logger.Info("getting menus by category", "categoryID", input.CategoryID)

	menus, err := u.menuRepo.GetByCategory(ctx, input.CategoryID)
	if err != nil {
		u.logger.Error("failed to get menus by category", "categoryID", input.CategoryID, "error", err)
		return nil, err
	}

	u.logger.Info("successfully retrieved menus by category", "categoryID", input.CategoryID, "count", len(menus))

	return &output.GetMenusByCategoryOutput{
		Menus:      menus,
		Count:      len(menus),
		CategoryID: input.CategoryID,
	}, nil
}

// CreateMenu メニューを作成する
func (u *MenuUsecaseImpl) CreateMenu(ctx context.Context, input *input.CreateMenuInput) (*output.CreateMenuOutput, error) {
	u.logger.Info("creating menu", "menuID", input.MenuID, "name", input.Name)

	// TODO: バリデーションを追加
	// TODO: トランザクション処理を追加
	// TODO: カテゴリの存在確認

	menuEntity := &menu.Menu{
		MenuID:      input.MenuID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		ImageURL:    input.ImageURL,
		IsSoldOut:   input.IsSoldOut,
		CategoryID:  input.CategoryID,
	}

	if err := u.menuRepo.Create(ctx, menuEntity); err != nil {
		u.logger.Error("failed to create menu", "menuID", input.MenuID, "error", err)
		return nil, err
	}

	// 作成されたメニューを取得
	createdMenu, err := u.menuRepo.GetByID(ctx, input.MenuID)
	if err != nil {
		u.logger.Error("failed to get created menu", "menuID", input.MenuID, "error", err)
		return nil, err
	}

	u.logger.Info("menu created successfully", "menuID", input.MenuID)

	return &output.CreateMenuOutput{
		Menu:        createdMenu,
		Message:     "Menu created successfully",
		CreatedAt:   time.Now(),
		MenuID:      createdMenu.MenuID,
		Name:        createdMenu.Name,
		Description: createdMenu.Description,
		Price:       createdMenu.Price,
		ImageURL:    createdMenu.ImageURL,
		IsSoldOut:   createdMenu.IsSoldOut,
		CategoryID:  createdMenu.CategoryID,
	}, nil
}

// UpdateMenu メニューを更新する
func (u *MenuUsecaseImpl) UpdateMenu(ctx context.Context, input *input.UpdateMenuInput) (*output.UpdateMenuOutput, error) {
	u.logger.Info("updating menu", "menuID", input.MenuID)

	// TODO: バリデーションを追加
	// TODO: トランザクション処理を追加
	// TODO: カテゴリの存在確認

	// 既存のメニューが存在するかチェック
	existingMenu, err := u.menuRepo.GetByID(ctx, input.MenuID)
	if err != nil {
		u.logger.Error("failed to get menu for update", "menuID", input.MenuID, "error", err)
		return nil, err
	}

	if existingMenu == nil {
		u.logger.Warn("menu not found for update", "menuID", input.MenuID)
		return nil, err // TODO: カスタムエラーに変更
	}

	menuEntity := &menu.Menu{
		MenuID:      input.MenuID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		ImageURL:    input.ImageURL,
		IsSoldOut:   input.IsSoldOut,
		CategoryID:  input.CategoryID,
	}

	if err := u.menuRepo.Update(ctx, menuEntity); err != nil {
		u.logger.Error("failed to update menu", "menuID", input.MenuID, "error", err)
		return nil, err
	}

	// 更新されたメニューを取得
	updatedMenu, err := u.menuRepo.GetByID(ctx, input.MenuID)
	if err != nil {
		u.logger.Error("failed to get updated menu", "menuID", input.MenuID, "error", err)
		return nil, err
	}

	u.logger.Info("menu updated successfully", "menuID", input.MenuID)

	return &output.UpdateMenuOutput{
		Menu:      updatedMenu,
		Message:   "Menu updated successfully",
		UpdatedAt: time.Now(),
	}, nil
}

// DeleteMenu メニューを削除する
func (u *MenuUsecaseImpl) DeleteMenu(ctx context.Context, input *input.DeleteMenuInput) (*output.DeleteMenuOutput, error) {
	u.logger.Info("deleting menu", "menuID", input.MenuID)

	// TODO: 関連データの存在チェック（注文履歴など）
	// TODO: トランザクション処理を追加

	// 既存のメニューが存在するかチェック
	existingMenu, err := u.menuRepo.GetByID(ctx, input.MenuID)
	if err != nil {
		u.logger.Error("failed to get menu for delete", "menuID", input.MenuID, "error", err)
		return nil, err
	}

	if existingMenu == nil {
		u.logger.Warn("menu not found for delete", "menuID", input.MenuID)
		return nil, err // TODO: カスタムエラーに変更
	}

	if err := u.menuRepo.Delete(ctx, input.MenuID); err != nil {
		u.logger.Error("failed to delete menu", "menuID", input.MenuID, "error", err)
		return nil, err
	}

	u.logger.Info("menu deleted successfully", "menuID", input.MenuID)

	return &output.DeleteMenuOutput{
		MenuID:    input.MenuID,
		Message:   "Menu deleted successfully",
		DeletedAt: time.Now(),
	}, nil
}
