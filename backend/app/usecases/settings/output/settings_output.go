package output

import (
	"time"

	"inoUwU/pinu/app/domain/settings"
)

// GetSettingsOutput 全設定取得の出力
type GetSettingsOutput struct {
	Settings []*settings.Settings `json:"settings"`
	Count    int                  `json:"count"`
}

// GetSettingByKeyOutput キー指定設定取得の出力
type GetSettingByKeyOutput struct {
	Setting *settings.Settings `json:"setting"`
	Found   bool               `json:"found"`
}

// SetSettingOutput 設定登録・更新の出力
type SetSettingOutput struct {
	Setting   *settings.Settings `json:"setting"`
	Message   string             `json:"message"`
	UpdatedAt time.Time          `json:"updated_at"`
	Key       string             `json:"key"`
	Value     string             `json:"value"`
	IsNew     bool               `json:"is_new"`
}

// DeleteSettingOutput 設定削除の出力
type DeleteSettingOutput struct {
	Key       string    `json:"key"`
	Message   string    `json:"message"`
	DeletedAt time.Time `json:"deleted_at"`
}
