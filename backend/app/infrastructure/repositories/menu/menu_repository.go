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
		MenuID:      menuItem.MenuID,
		Name:        menuItem.Name,
		Description: menuItem.Description,
		Price:       menuItem.Price,
		ImageURL:    menuItem.ImageURL,
		IsSoldOut:   menuItem.IsSoldOut,
		CategoryID:  menuItem.CategoryID,
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
		MenuID:      model.MenuID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		ImageURL:    model.ImageURL,
		IsSoldOut:   model.IsSoldOut,
		CategoryID:  model.CategoryID,
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
			MenuID:      model.MenuID,
			Name:        model.Name,
			Description: model.Description,
			Price:       model.Price,
			ImageURL:    model.ImageURL,
			IsSoldOut:   model.IsSoldOut,
			CategoryID:  model.CategoryID,
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
			MenuID:      model.MenuID,
			Name:        model.Name,
			Description: model.Description,
			Price:       model.Price,
			ImageURL:    model.ImageURL,
			IsSoldOut:   model.IsSoldOut,
			CategoryID:  model.CategoryID,
		}
	}

	return menus, nil
}

func (r *menuRepository) Update(ctx context.Context, menuItem *menu.Menu) error {
	model := &models.MenuModel{
		MenuID:      menuItem.MenuID,
		Name:        menuItem.Name,
		Description: menuItem.Description,
		Price:       menuItem.Price,
		ImageURL:    menuItem.ImageURL,
		IsSoldOut:   menuItem.IsSoldOut,
		CategoryID:  menuItem.CategoryID,
	}

	_, err := r.db.NewUpdate().
		Model(model).
		Where("menu_id = ?", menuItem.MenuID).
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
