package repositories

import (
	"server/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	FindAllUser() ([]model.User, error)
	GetUser(ID int) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(user model.User) (model.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) FindAllUser() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetUser(ID int) (model.User, error) {
	var user model.User
	err := r.db.First(&user, ID).Error
	return user, err
}

func (r *repository) UpdateUser(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user model.User) (model.User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
