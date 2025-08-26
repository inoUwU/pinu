package settings

import (
	"context"
	"time"

	"github.com/samber/do"

	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/domain/settings"
	"inoUwU/pinu/app/usecases/settings/input"
	"inoUwU/pinu/app/usecases/settings/output"
)

// ISettingsUsecase 設定ユースケースのインターフェース
type ISettingsUsecase interface {
	GetAllSettings(ctx context.Context, input *input.GetSettingsInput) (*output.GetSettingsOutput, error)
	GetSettingByKey(ctx context.Context, input *input.GetSettingByKeyInput) (*output.GetSettingByKeyOutput, error)
	SetSetting(ctx context.Context, input *input.SetSettingInput) (*output.SetSettingOutput, error)
	DeleteSetting(ctx context.Context, input *input.DeleteSettingInput) (*output.DeleteSettingOutput, error)
}

// SettingsUsecaseImpl 設定ユースケースの実装
type SettingsUsecaseImpl struct {
	settingsRepo settings.SettingsRepository
	logger       port.Logger
}

// NewSettingsUsecase 設定ユースケースを生成する
func NewSettingsUsecase(i *do.Injector) (ISettingsUsecase, error) {
	repository := do.MustInvoke[settings.SettingsRepository](i)
	logger := do.MustInvokeNamed[port.Logger](i, "logger")

	return &SettingsUsecaseImpl{
		settingsRepo: repository,
		logger:       logger,
	}, nil
}

// GetAllSettings 全ての設定を取得する
func (u *SettingsUsecaseImpl) GetAllSettings(ctx context.Context, input *input.GetSettingsInput) (*output.GetSettingsOutput, error) {
	u.logger.Info("getting all settings")

	settingsList, err := u.settingsRepo.GetAll(ctx)
	if err != nil {
		u.logger.Error("failed to get all settings", "error", err)
		return nil, err
	}

	u.logger.Info("successfully retrieved settings", "count", len(settingsList))

	return &output.GetSettingsOutput{
		Settings: settingsList,
		Count:    len(settingsList),
	}, nil
}

// GetSettingByKey 指定キーの設定を取得する
func (u *SettingsUsecaseImpl) GetSettingByKey(ctx context.Context, input *input.GetSettingByKeyInput) (*output.GetSettingByKeyOutput, error) {
	u.logger.Info("getting setting by key", "key", input.Key)

	setting, err := u.settingsRepo.Get(ctx, input.Key)
	if err != nil {
		u.logger.Error("failed to get setting by key", "key", input.Key, "error", err)
		return nil, err
	}

	found := setting != nil
	if !found {
		u.logger.Warn("setting not found", "key", input.Key)
	} else {
		u.logger.Info("successfully retrieved setting", "key", input.Key)
	}

	return &output.GetSettingByKeyOutput{
		Setting: setting,
		Found:   found,
	}, nil
}

// SetSetting 設定を登録・更新する
func (u *SettingsUsecaseImpl) SetSetting(ctx context.Context, input *input.SetSettingInput) (*output.SetSettingOutput, error) {
	u.logger.Info("setting configuration", "key", input.Key)

	// TODO: バリデーションを追加
	// TODO: トランザクション処理を追加

	// 既存の設定があるかチェック
	existingSetting, err := u.settingsRepo.Get(ctx, input.Key)
	if err != nil {
		u.logger.Error("failed to check existing setting", "key", input.Key, "error", err)
		return nil, err
	}

	isNew := existingSetting == nil

	if err := u.settingsRepo.Set(ctx, input.Key, input.Value); err != nil {
		u.logger.Error("failed to set setting", "key", input.Key, "error", err)
		return nil, err
	}

	// 更新された設定を取得
	updatedSetting, err := u.settingsRepo.Get(ctx, input.Key)
	if err != nil {
		u.logger.Error("failed to get updated setting", "key", input.Key, "error", err)
		return nil, err
	}

	message := "Setting updated successfully"
	if isNew {
		message = "Setting created successfully"
	}

	u.logger.Info("setting operation completed", "key", input.Key, "isNew", isNew)

	return &output.SetSettingOutput{
		Setting:   updatedSetting,
		Message:   message,
		UpdatedAt: time.Now(),
		Key:       input.Key,
		Value:     input.Value,
		IsNew:     isNew,
	}, nil
}

// DeleteSetting 設定を削除する
func (u *SettingsUsecaseImpl) DeleteSetting(ctx context.Context, input *input.DeleteSettingInput) (*output.DeleteSettingOutput, error) {
	u.logger.Info("deleting setting", "key", input.Key)

	// TODO: 重要な設定の削除防止機能を追加
	// TODO: トランザクション処理を追加

	// 既存の設定が存在するかチェック
	existingSetting, err := u.settingsRepo.Get(ctx, input.Key)
	if err != nil {
		u.logger.Error("failed to get setting for delete", "key", input.Key, "error", err)
		return nil, err
	}

	if existingSetting == nil {
		u.logger.Warn("setting not found for delete", "key", input.Key)
		return nil, err // TODO: カスタムエラーに変更
	}

	if err := u.settingsRepo.Delete(ctx, input.Key); err != nil {
		u.logger.Error("failed to delete setting", "key", input.Key, "error", err)
		return nil, err
	}

	u.logger.Info("setting deleted successfully", "key", input.Key)

	return &output.DeleteSettingOutput{
		Key:       input.Key,
		Message:   "Setting deleted successfully",
		DeletedAt: time.Now(),
	}, nil
}
