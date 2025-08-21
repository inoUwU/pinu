package category

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"

	"inoUwU/pinu/app/domain/category"
	"inoUwU/pinu/app/infrastructure/models"
)

type categoryRepository struct {
	db *bun.DB
}

// NewCategoryRepository カテゴリリポジトリの新規作成
func NewCategoryRepository(db *bun.DB) category.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx context.Context, cat *category.Category) error {
	model := &models.CategoryModel{
		CategoryID:   cat.CategoryID,
		Name:         cat.Name,
		DisplayOrder: cat.DisplayOrder,
		ImageURL:     cat.ImageURL,
	}

	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	return err
}

func (r *categoryRepository) GetByID(ctx context.Context, id string) (*category.Category, error) {
	model := &models.CategoryModel{}
	err := r.db.NewSelect().
		Model(model).
		Where("category_id = ?", id).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &category.Category{
		CategoryID:   model.CategoryID,
		Name:         model.Name,
		DisplayOrder: model.DisplayOrder,
		ImageURL:     model.ImageURL,
	}, nil
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]*category.Category, error) {
	var models []*models.CategoryModel
	err := r.db.NewSelect().
		Model(&models).
		Order("display_order ASC, category_id ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	categories := make([]*category.Category, len(models))
	for i, model := range models {
		categories[i] = &category.Category{
			CategoryID:   model.CategoryID,
			Name:         model.Name,
			DisplayOrder: model.DisplayOrder,
			ImageURL:     model.ImageURL,
		}
	}

	return categories, nil
}

func (r *categoryRepository) Update(ctx context.Context, cat *category.Category) error {
	model := &models.CategoryModel{
		CategoryID:   cat.CategoryID,
		Name:         cat.Name,
		DisplayOrder: cat.DisplayOrder,
		ImageURL:     cat.ImageURL,
	}

	_, err := r.db.NewUpdate().
		Model(model).
		Where("category_id = ?", cat.CategoryID).
		Exec(ctx)

	return err
}

func (r *categoryRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewDelete().
		Model((*models.CategoryModel)(nil)).
		Where("category_id = ?", id).
		Exec(ctx)

	return err
}
