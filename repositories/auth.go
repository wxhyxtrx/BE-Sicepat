package repositories

import (
	"server/model"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user model.User) (model.User, error)
	Login(email string) ([]model.User, error)
	CekUser(email string) (model.User, error)
	UserLogin(id int) (model.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(email string) ([]model.User, error) {
	var user []model.User
	err := r.db.Find(&user, "email=?", email).Error

	return user, err
}
func (r *repository) UserLogin(id int) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return user, err
}
func (r *repository) CekUser(email string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, "email=? AND status = 'AKTIF'", email).Error
	return user, err
}
