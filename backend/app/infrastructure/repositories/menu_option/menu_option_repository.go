package menu_option

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"

	"inoUwU/pinu/app/domain/menu_option"
	"inoUwU/pinu/app/infrastructure/models"
)

type menuOptionRepository struct {
	db *bun.DB
}

// NewMenuOptionRepository メニューオプションリポジトリの新規作成
func NewMenuOptionRepository(db *bun.DB) menu_option.MenuOptionRepository {
	return &menuOptionRepository{db: db}
}

func (r *menuOptionRepository) Create(ctx context.Context, menuOption *menu_option.MenuOption) error {
	model := &models.MenuOptionModel{
		MenuOptionID: string(menuOption.MenuOptionID),
		Name:         menuOption.Name,
		Price:        menuOption.Price,
	}

	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	return err
}

func (r *menuOptionRepository) GetByID(ctx context.Context, id menu_option.MenuOptionID) (*menu_option.MenuOption, error) {
	model := &models.MenuOptionModel{}
	err := r.db.NewSelect().
		Model(model).
		Where("menu_option_id = ?", string(id)).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &menu_option.MenuOption{
		MenuOptionID: menu_option.MenuOptionID(model.MenuOptionID),
		Name:         model.Name,
		Price:        model.Price,
	}, nil
}

func (r *menuOptionRepository) GetAll(ctx context.Context) ([]*menu_option.MenuOption, error) {
	var models []*models.MenuOptionModel
	err := r.db.NewSelect().
		Model(&models).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	menuOptions := make([]*menu_option.MenuOption, len(models))
	for i, model := range models {
		menuOptions[i] = &menu_option.MenuOption{
			MenuOptionID: menu_option.MenuOptionID(model.MenuOptionID),
			Name:         model.Name,
			Price:        model.Price,
		}
	}

	return menuOptions, nil
}

func (r *menuOptionRepository) Update(ctx context.Context, menuOption *menu_option.MenuOption) error {
	model := &models.MenuOptionModel{
		MenuOptionID: string(menuOption.MenuOptionID),
		Name:         menuOption.Name,
		Price:        menuOption.Price,
	}

	_, err := r.db.NewUpdate().
		Model(model).
		Where("menu_option_id = ?", string(menuOption.MenuOptionID)).
		Exec(ctx)

	return err
}

func (r *menuOptionRepository) Delete(ctx context.Context, id menu_option.MenuOptionID) error {
	_, err := r.db.NewDelete().
		Model((*models.MenuOptionModel)(nil)).
		Where("menu_option_id = ?", string(id)).
		Exec(ctx)

	return err
}
