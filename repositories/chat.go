package repositories

import (
	"server/model"

	"gorm.io/gorm"
)

type ChatRepository interface {
	CreateChat(chat model.Chat) (model.Chat, error)
	GetChat(id int) (model.Chat, error)
	CheckSetting() ([]model.Setting, error)

	User(ID int) (model.User, error)
	SettingUser(user model.User) (model.User, error)
}

func RepositoryChat(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateChat(chat model.Chat) (model.Chat, error) {
	err := r.db.Create(&chat).Error
	return chat, err
}

func (r *repository) GetChat(id int) (model.Chat, error) {
	var chat model.Chat
	err := r.db.Preload("From").Preload("Room").First(&chat, id).Error
	return chat, err
}

// setting
func (r *repository) CheckSetting() ([]model.Setting, error) {
	var setting []model.Setting
	err := r.db.Find(&setting).Error

	return setting, err
}

// User
func (r *repository) User(ID int) (model.User, error) {
	var user model.User
	err := r.db.First(&user, ID).Error
	return user, err
}

func (r *repository) SettingUser(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}
