package category

import (
	"context"
	"time"

	"github.com/samber/do"

	"inoUwU/pinu/app/domain/category"
	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/usecases/category/input"
	"inoUwU/pinu/app/usecases/category/output"
)

// CategoryService カテゴリユースケースのポート
type CategoryService interface {
	GetAllCategories(ctx context.Context, input *input.GetCategoriesInput) (*output.GetCategoriesOutput, error)
	GetCategoryByID(ctx context.Context, input *input.GetCategoryByIDInput) (*output.GetCategoryByIDOutput, error)
	CreateCategory(ctx context.Context, input *input.CreateCategoryInput) (*output.CreateCategoryOutput, error)
	UpdateCategory(ctx context.Context, input *input.UpdateCategoryInput) (*output.UpdateCategoryOutput, error)
	DeleteCategory(ctx context.Context, input *input.DeleteCategoryInput) (*output.DeleteCategoryOutput, error)
}

// CategoryUsecaseImpl カテゴリユースケースの実装
type CategoryUsecaseImpl struct {
	categoryRepo category.CategoryRepository
	logger       port.Logger
}

// NewCategoryUsecase カテゴリユースケースを生成する
func NewCategoryUsecase(i *do.Injector) (CategoryService, error) {
	repository := do.MustInvoke[category.CategoryRepository](i)
	logger := do.MustInvokeNamed[port.Logger](i, "logger")

	return &CategoryUsecaseImpl{
		categoryRepo: repository,
		logger:       logger,
	}, nil
}

// GetAllCategories 全てのカテゴリを取得する
func (u *CategoryUsecaseImpl) GetAllCategories(ctx context.Context, input *input.GetCategoriesInput) (*output.GetCategoriesOutput, error) {
	u.logger.Info("getting all categories")

	categories, err := u.categoryRepo.GetAll(ctx)
	if err != nil {
		u.logger.Error("failed to get all categories", "error", err)
		return nil, err
	}

	u.logger.Info("successfully retrieved categories", "count", len(categories))

	return &output.GetCategoriesOutput{
		Categories: categories,
		Count:      len(categories),
	}, nil
}

// GetCategoryByID 指定IDのカテゴリを取得する
func (u *CategoryUsecaseImpl) GetCategoryByID(ctx context.Context, input *input.GetCategoryByIDInput) (*output.GetCategoryByIDOutput, error) {
	u.logger.Info("getting category by ID", "categoryID", input.CategoryID)

	category, err := u.categoryRepo.GetByID(ctx, input.CategoryID)
	if err != nil {
		u.logger.Error("failed to get category by ID", "categoryID", input.CategoryID, "error", err)
		return nil, err
	}

	if category == nil {
		u.logger.Warn("category not found", "categoryID", input.CategoryID)
		return &output.GetCategoryByIDOutput{
			Category: nil,
		}, nil
	}

	u.logger.Info("successfully retrieved category", "categoryID", input.CategoryID)

	return &output.GetCategoryByIDOutput{
		Category: category,
	}, nil
}

// CreateCategory カテゴリを作成する
func (u *CategoryUsecaseImpl) CreateCategory(ctx context.Context, input *input.CreateCategoryInput) (*output.CreateCategoryOutput, error) {
	u.logger.Info("creating category", "categoryID", input.CategoryID, "name", input.Name)

	// TODO: バリデーションを追加
	// TODO: トランザクション処理を追加

	categoryEntity := &category.Category{
		CategoryID:   input.CategoryID,
		Name:         input.Name,
		DisplayOrder: input.DisplayOrder,
		ImageURL:     input.ImageURL,
	}

	if err := u.categoryRepo.Create(ctx, categoryEntity); err != nil {
		u.logger.Error("failed to create category", "categoryID", input.CategoryID, "error", err)
		return nil, err
	}

	// 作成されたカテゴリを取得
	createdCategory, err := u.categoryRepo.GetByID(ctx, input.CategoryID)
	if err != nil {
		u.logger.Error("failed to get created category", "categoryID", input.CategoryID, "error", err)
		return nil, err
	}

	u.logger.Info("category created successfully", "categoryID", input.CategoryID)

	return &output.CreateCategoryOutput{
		Category:     createdCategory,
		Message:      "Category created successfully",
		CreatedAt:    time.Now(),
		CategoryID:   createdCategory.CategoryID,
		Name:         createdCategory.Name,
		DisplayOrder: createdCategory.DisplayOrder,
		ImageURL:     createdCategory.ImageURL,
	}, nil
}

// UpdateCategory カテゴリを更新する
func (u *CategoryUsecaseImpl) UpdateCategory(ctx context.Context, input *input.UpdateCategoryInput) (*output.UpdateCategoryOutput, error) {
	u.logger.Info("updating category", "categoryID", input.CategoryID)

	// TODO: バリデーションを追加
	// TODO: トランザクション処理を追加

	// 既存のカテゴリが存在するかチェック
	existingCategory, err := u.categoryRepo.GetByID(ctx, input.CategoryID)
	if err != nil {
		u.logger.Error("failed to get category for update", "categoryID", input.CategoryID, "error", err)
		return nil, err
	}

	if existingCategory == nil {
		u.logger.Warn("category not found for update", "categoryID", input.CategoryID)
		return nil, err // TODO: カスタムエラーに変更
	}

	categoryEntity := &category.Category{
		CategoryID:   input.CategoryID,
		Name:         input.Name,
		DisplayOrder: input.DisplayOrder,
		ImageURL:     input.ImageURL,
	}

	if err := u.categoryRepo.Update(ctx, categoryEntity); err != nil {
		u.logger.Error("failed to update category", "categoryID", input.CategoryID, "error", err)
		return nil, err
	}

	// 更新されたカテゴリを取得
	updatedCategory, err := u.categoryRepo.GetByID(ctx, input.CategoryID)
	if err != nil {
		u.logger.Error("failed to get updated category", "categoryID", input.CategoryID, "error", err)
		return nil, err
	}

	u.logger.Info("category updated successfully", "categoryID", input.CategoryID)

	return &output.UpdateCategoryOutput{
		Category:  updatedCategory,
		Message:   "Category updated successfully",
		UpdatedAt: time.Now(),
	}, nil
}

// DeleteCategory カテゴリを削除する
func (u *CategoryUsecaseImpl) DeleteCategory(ctx context.Context, input *input.DeleteCategoryInput) (*output.DeleteCategoryOutput, error) {
	u.logger.Info("deleting category", "categoryID", input.CategoryID)

	// TODO: 関連データの存在チェック
	// TODO: トランザクション処理を追加

	// 既存のカテゴリが存在するかチェック
	existingCategory, err := u.categoryRepo.GetByID(ctx, input.CategoryID)
	if err != nil {
		u.logger.Error("failed to get category for delete", "categoryID", input.CategoryID, "error", err)
		return nil, err
	}

	if existingCategory == nil {
		u.logger.Warn("category not found for delete", "categoryID", input.CategoryID)
		return nil, err // TODO: カスタムエラーに変更
	}

	if err := u.categoryRepo.Delete(ctx, input.CategoryID); err != nil {
		u.logger.Error("failed to delete category", "categoryID", input.CategoryID, "error", err)
		return nil, err
	}

	u.logger.Info("category deleted successfully", "categoryID", input.CategoryID)

	return &output.DeleteCategoryOutput{
		CategoryID: input.CategoryID,
		Message:    "Category deleted successfully",
		DeletedAt:  time.Now(),
	}, nil
}
