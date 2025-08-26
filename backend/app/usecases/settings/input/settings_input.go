package input

// GetSettingsInput 全設定取得の入力
type GetSettingsInput struct{}

// GetSettingByKeyInput キー指定設定取得の入力
type GetSettingByKeyInput struct {
	Key string `json:"key" validate:"required"`
}

// SetSettingInput 設定登録・更新の入力
type SetSettingInput struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// DeleteSettingInput 設定削除の入力
type DeleteSettingInput struct {
	Key string `json:"key" validate:"required"`
}
