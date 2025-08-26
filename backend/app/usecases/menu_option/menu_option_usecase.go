package menu_option

import (
	"context"
	"time"

	"github.com/samber/do"

	"inoUwU/pinu/app/domain/menu_option"
	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/usecases/menu_option/input"
	"inoUwU/pinu/app/usecases/menu_option/output"
)

// IMenuOptionUsecase メニューオプションユースケースのインターフェース
type IMenuOptionUsecase interface {
	GetAllMenuOptions(ctx context.Context, input *input.GetMenuOptionsInput) (*output.GetMenuOptionsOutput, error)
	GetMenuOptionByID(ctx context.Context, input *input.GetMenuOptionByIDInput) (*output.GetMenuOptionByIDOutput, error)
	CreateMenuOption(ctx context.Context, input *input.CreateMenuOptionInput) (*output.CreateMenuOptionOutput, error)
	UpdateMenuOption(ctx context.Context, input *input.UpdateMenuOptionInput) (*output.UpdateMenuOptionOutput, error)
	DeleteMenuOption(ctx context.Context, input *input.DeleteMenuOptionInput) (*output.DeleteMenuOptionOutput, error)
}

// MenuOptionUsecaseImpl メニューオプションユースケースの実装
type MenuOptionUsecaseImpl struct {
	menuOptionRepo menu_option.MenuOptionRepository
	logger         port.Logger
}

// NewMenuOptionUsecase メニューオプションユースケースを生成する
func NewMenuOptionUsecase(i *do.Injector) (IMenuOptionUsecase, error) {
	repository := do.MustInvoke[menu_option.MenuOptionRepository](i)
	logger := do.MustInvokeNamed[port.Logger](i, "logger")

	return &MenuOptionUsecaseImpl{
		menuOptionRepo: repository,
		logger:         logger,
	}, nil
}

// GetAllMenuOptions 全てのメニューオプションを取得する
func (u *MenuOptionUsecaseImpl) GetAllMenuOptions(ctx context.Context, input *input.GetMenuOptionsInput) (*output.GetMenuOptionsOutput, error) {
	u.logger.Info("getting all menu options")

	menuOptions, err := u.menuOptionRepo.GetAll(ctx)
	if err != nil {
		u.logger.Error("failed to get all menu options", "error", err)
		return nil, err
	}

	u.logger.Info("successfully retrieved menu options", "count", len(menuOptions))

	return &output.GetMenuOptionsOutput{
		MenuOptions: menuOptions,
		Count:       len(menuOptions),
	}, nil
}

// GetMenuOptionByID 指定IDのメニューオプションを取得する
func (u *MenuOptionUsecaseImpl) GetMenuOptionByID(ctx context.Context, input *input.GetMenuOptionByIDInput) (*output.GetMenuOptionByIDOutput, error) {
	u.logger.Info("getting menu option by ID", "menuOptionID", input.MenuOptionID)

	menuOption, err := u.menuOptionRepo.GetByID(ctx, input.MenuOptionID)
	if err != nil {
		u.logger.Error("failed to get menu option by ID", "menuOptionID", input.MenuOptionID, "error", err)
		return nil, err
	}

	if menuOption == nil {
		u.logger.Warn("menu option not found", "menuOptionID", input.MenuOptionID)
		return &output.GetMenuOptionByIDOutput{
			MenuOption: nil,
		}, nil
	}

	u.logger.Info("successfully retrieved menu option", "menuOptionID", input.MenuOptionID)

	return &output.GetMenuOptionByIDOutput{
		MenuOption: menuOption,
	}, nil
}

// CreateMenuOption メニューオプションを作成する
func (u *MenuOptionUsecaseImpl) CreateMenuOption(ctx context.Context, input *input.CreateMenuOptionInput) (*output.CreateMenuOptionOutput, error) {
	u.logger.Info("creating menu option", "menuOptionID", input.MenuOptionID, "name", input.Name)

	// TODO: バリデーションを追加
	// TODO: トランザクション処理を追加

	menuOptionEntity := &menu_option.MenuOption{
		MenuOptionID: input.MenuOptionID,
		Name:         input.Name,
		Price:        input.Price,
	}

	if err := u.menuOptionRepo.Create(ctx, menuOptionEntity); err != nil {
		u.logger.Error("failed to create menu option", "menuOptionID", input.MenuOptionID, "error", err)
		return nil, err
	}

	// 作成されたメニューオプションを取得
	createdMenuOption, err := u.menuOptionRepo.GetByID(ctx, input.MenuOptionID)
	if err != nil {
		u.logger.Error("failed to get created menu option", "menuOptionID", input.MenuOptionID, "error", err)
		return nil, err
	}

	u.logger.Info("menu option created successfully", "menuOptionID", input.MenuOptionID)

	return &output.CreateMenuOptionOutput{
		MenuOption:   createdMenuOption,
		Message:      "Menu option created successfully",
		CreatedAt:    time.Now(),
		MenuOptionID: createdMenuOption.MenuOptionID,
		Name:         createdMenuOption.Name,
		Price:        createdMenuOption.Price,
	}, nil
}

// UpdateMenuOption メニューオプションを更新する
func (u *MenuOptionUsecaseImpl) UpdateMenuOption(ctx context.Context, input *input.UpdateMenuOptionInput) (*output.UpdateMenuOptionOutput, error) {
	u.logger.Info("updating menu option", "menuOptionID", input.MenuOptionID)

	// TODO: バリデーションを追加
	// TODO: トランザクション処理を追加

	// 既存のメニューオプションが存在するかチェック
	existingMenuOption, err := u.menuOptionRepo.GetByID(ctx, input.MenuOptionID)
	if err != nil {
		u.logger.Error("failed to get menu option for update", "menuOptionID", input.MenuOptionID, "error", err)
		return nil, err
	}

	if existingMenuOption == nil {
		u.logger.Warn("menu option not found for update", "menuOptionID", input.MenuOptionID)
		return nil, err // TODO: カスタムエラーに変更
	}

	menuOptionEntity := &menu_option.MenuOption{
		MenuOptionID: input.MenuOptionID,
		Name:         input.Name,
		Price:        input.Price,
	}

	if err := u.menuOptionRepo.Update(ctx, menuOptionEntity); err != nil {
		u.logger.Error("failed to update menu option", "menuOptionID", input.MenuOptionID, "error", err)
		return nil, err
	}

	// 更新されたメニューオプションを取得
	updatedMenuOption, err := u.menuOptionRepo.GetByID(ctx, input.MenuOptionID)
	if err != nil {
		u.logger.Error("failed to get updated menu option", "menuOptionID", input.MenuOptionID, "error", err)
		return nil, err
	}

	u.logger.Info("menu option updated successfully", "menuOptionID", input.MenuOptionID)

	return &output.UpdateMenuOptionOutput{
		MenuOption: updatedMenuOption,
		Message:    "Menu option updated successfully",
		UpdatedAt:  time.Now(),
	}, nil
}

// DeleteMenuOption メニューオプションを削除する
func (u *MenuOptionUsecaseImpl) DeleteMenuOption(ctx context.Context, input *input.DeleteMenuOptionInput) (*output.DeleteMenuOptionOutput, error) {
	u.logger.Info("deleting menu option", "menuOptionID", input.MenuOptionID)

	// TODO: 関連データの存在チェック
	// TODO: トランザクション処理を追加

	// 既存のメニューオプションが存在するかチェック
	existingMenuOption, err := u.menuOptionRepo.GetByID(ctx, input.MenuOptionID)
	if err != nil {
		u.logger.Error("failed to get menu option for delete", "menuOptionID", input.MenuOptionID, "error", err)
		return nil, err
	}

	if existingMenuOption == nil {
		u.logger.Warn("menu option not found for delete", "menuOptionID", input.MenuOptionID)
		return nil, err // TODO: カスタムエラーに変更
	}

	if err := u.menuOptionRepo.Delete(ctx, input.MenuOptionID); err != nil {
		u.logger.Error("failed to delete menu option", "menuOptionID", input.MenuOptionID, "error", err)
		return nil, err
	}

	u.logger.Info("menu option deleted successfully", "menuOptionID", input.MenuOptionID)

	return &output.DeleteMenuOptionOutput{
		MenuOptionID: input.MenuOptionID,
		Message:      "Menu option deleted successfully",
		DeletedAt:    time.Now(),
	}, nil
}
