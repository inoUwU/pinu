package models

import (
	"github.com/uptrace/bun"
)

// SettingsModel 設定のインフラストラクチャーモデル
type SettingsModel struct {
	bun.BaseModel `bun:"table:settings"`

	Key   string `bun:"key,pk"`
	Value string `bun:"value,notnull"`
}
