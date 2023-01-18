package repositories

import (
	"server/model"

	"gorm.io/gorm"
)

type TiketRepository interface {
	CreateTiket(tiket model.Tiket) (model.Tiket, error)
	FindTiket() ([]model.Tiket, error)
	GetTiket(ID int) (model.Tiket, error)
}

func RepositoryTiket(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTiket(tiket model.Tiket) (model.Tiket, error) {
	err := r.db.Create(&tiket).Error
	return tiket, err
}
func (r *repository) FindTiket() ([]model.Tiket, error) {
	var tikets []model.Tiket
	err := r.db.Find(&tikets).Error
	return tikets, err
}

func (r *repository) GetTiket(ID int) (model.Tiket, error) {
	var tiket model.Tiket
	err := r.db.First(&tiket, ID).Error
	return tiket, err
}
