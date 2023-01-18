package repositories

import (
	"server/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(item model.Order) (model.Order, error)
	GetOrder(id int) (model.Order, error)

	CheckTiket(id int) (model.Tiket, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateOrder(item model.Order) (model.Order, error) {
	err := r.db.Create(&item).Error

	var tiket model.Tiket
	r.db.First(&tiket, item.TiketID)
	tiket.Stok = tiket.Stok - item.Qty
	r.db.Save(&tiket)

	return item, err
}

func (r *repository) GetOrder(id int) (model.Order, error) {
	var item model.Order
	err := r.db.Preload("Tiket").Preload("User").First(&item, id).Error

	return item, err
}

// tiket

func (r *repository) CheckTiket(id int) (model.Tiket, error) {
	var tiket model.Tiket
	err := r.db.First(&tiket, id).Error

	return tiket, err
}
