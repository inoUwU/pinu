package settings

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"

	"inoUwU/pinu/app/domain/settings"
	"inoUwU/pinu/app/infrastructure/models"
)

type settingsRepository struct {
	db *bun.DB
}

// NewSettingsRepository 設定リポジトリの新規作成
func NewSettingsRepository(db *bun.DB) settings.SettingsRepository {
	return &settingsRepository{db: db}
}

func (r *settingsRepository) Get(ctx context.Context, key string) (*settings.Settings, error) {
	model := &models.SettingsModel{}
	err := r.db.NewSelect().
		Model(model).
		Where("key = ?", key).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &settings.Settings{
		Key:   model.Key,
		Value: model.Value,
	}, nil
}

func (r *settingsRepository) GetAll(ctx context.Context) ([]*settings.Settings, error) {
	var models []*models.SettingsModel
	err := r.db.NewSelect().
		Model(&models).
		Order("key ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	settingsList := make([]*settings.Settings, len(models))
	for i, model := range models {
		settingsList[i] = &settings.Settings{
			Key:   model.Key,
			Value: model.Value,
		}
	}

	return settingsList, nil
}

func (r *settingsRepository) Set(ctx context.Context, key, value string) error {
	model := &models.SettingsModel{
		Key:   key,
		Value: value,
	}

	_, err := r.db.NewInsert().
		Model(model).
		On("CONFLICT (key) DO UPDATE").
		Set("value = EXCLUDED.value").
		Exec(ctx)

	return err
}

func (r *settingsRepository) Delete(ctx context.Context, key string) error {
	_, err := r.db.NewDelete().
		Model((*models.SettingsModel)(nil)).
		Where("key = ?", key).
		Exec(ctx)

	return err
}
