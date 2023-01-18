package repositories

import (
	"server/model"

	"gorm.io/gorm"
)

type SettingRepository interface {
	CreateSetting(setting model.Setting) (model.Setting, error)
	FindSetting() ([]model.Setting, error)
	GetSetting(id int) (model.Setting, error)
}

func RepositorySetting(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateSetting(setting model.Setting) (model.Setting, error) {
	err := r.db.Create(&setting).Error
	return setting, err
}

func (r *repository) FindSetting() ([]model.Setting, error) {
	var setting []model.Setting
	err := r.db.Create(&setting).Error
	return setting, err
}

func (r *repository) GetSetting(id int) (model.Setting, error) {
	var setting model.Setting
	err := r.db.First(&setting, id).Error
	return setting, err
}
