package settings

import (
	"context"
)

// SettingsRepository 設定リポジトリのインターface
type SettingsRepository interface {
	Get(ctx context.Context, key string) (*Settings, error)
	GetAll(ctx context.Context) ([]*Settings, error)
	Set(ctx context.Context, key, value string) error
	Delete(ctx context.Context, key string) error
}
