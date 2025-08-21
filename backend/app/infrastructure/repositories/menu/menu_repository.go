package menu

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"

	"inoUwU/pinu/app/domain/menu"
	"inoUwU/pinu/app/infrastructure/models"
)

// menuRepository メニューリポジトリの実装
type menuRepository struct {
	db *bun.DB
}

// NewMenuRepository メニューリポジトリの新規作成
func NewMenuRepository(db *bun.DB) menu.MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) Create(ctx context.Context, menuItem *menu.Menu) error {
	model := &models.MenuModel{
		MenuID:      menuItem.MENU_ID,
		Name:        menuItem.NAME,
		Description: menuItem.DESCRIPTION,
		Price:       menuItem.PRICE,
		ImageURL:    menuItem.IMAGE_URL,
		IsSoldOut:   menuItem.IS_SOLD_OUT,
		CategoryID:  menuItem.CATEGORY_ID,
	}

	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	return err
}

func (r *menuRepository) GetByID(ctx context.Context, id string) (*menu.Menu, error) {
	model := &models.MenuModel{}
	err := r.db.NewSelect().
		Model(model).
		Where("menu_id = ?", id).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &menu.Menu{
		MENU_ID:     model.MenuID,
		NAME:        model.Name,
		DESCRIPTION: model.Description,
		PRICE:       model.Price,
		IMAGE_URL:   model.ImageURL,
		IS_SOLD_OUT: model.IsSoldOut,
		CATEGORY_ID: model.CategoryID,
	}, nil
}

func (r *menuRepository) GetAll(ctx context.Context) ([]*menu.Menu, error) {
	var models []*models.MenuModel
	err := r.db.NewSelect().
		Model(&models).
		Order("menu_id ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	menus := make([]*menu.Menu, len(models))
	for i, model := range models {
		menus[i] = &menu.Menu{
			MENU_ID:     model.MenuID,
			NAME:        model.Name,
			DESCRIPTION: model.Description,
			PRICE:       model.Price,
			IMAGE_URL:   model.ImageURL,
			IS_SOLD_OUT: model.IsSoldOut,
			CATEGORY_ID: model.CategoryID,
		}
	}

	return menus, nil
}

func (r *menuRepository) GetByCategory(ctx context.Context, categoryID string) ([]*menu.Menu, error) {
	var models []*models.MenuModel
	err := r.db.NewSelect().
		Model(&models).
		Where("category_id = ?", categoryID).
		Order("menu_id ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	menus := make([]*menu.Menu, len(models))
	for i, model := range models {
		menus[i] = &menu.Menu{
			MENU_ID:     model.MenuID,
			NAME:        model.Name,
			DESCRIPTION: model.Description,
			PRICE:       model.Price,
			IMAGE_URL:   model.ImageURL,
			IS_SOLD_OUT: model.IsSoldOut,
			CATEGORY_ID: model.CategoryID,
		}
	}

	return menus, nil
}

func (r *menuRepository) Update(ctx context.Context, menuItem *menu.Menu) error {
	model := &models.MenuModel{
		MenuID:      menuItem.MENU_ID,
		Name:        menuItem.NAME,
		Description: menuItem.DESCRIPTION,
		Price:       menuItem.PRICE,
		ImageURL:    menuItem.IMAGE_URL,
		IsSoldOut:   menuItem.IS_SOLD_OUT,
		CategoryID:  menuItem.CATEGORY_ID,
	}

	_, err := r.db.NewUpdate().
		Model(model).
		Where("menu_id = ?", menuItem.MENU_ID).
		Exec(ctx)

	return err
}

func (r *menuRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewDelete().
		Model((*models.MenuModel)(nil)).
		Where("menu_id = ?", id).
		Exec(ctx)

	return err
}
