package repositories

import (
	"server/model"

	"gorm.io/gorm"
)

type RoomRepository interface {
	// Deffault CRUD
	CreateRoom(room model.Room) (model.Room, error)
	FindsRoom() ([]model.Room, error)
	GetRoom(id int) (model.Room, error)
	UpdateRoom(room model.Room) (model.Room, error)
	DeleteRoom(room model.Room) (model.Room, error)
}

func RepositoryRoom(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateRoom(room model.Room) (model.Room, error) {
	err := r.db.Preload("user").Create(&room).Error
	return room, err
}

func (r *repository) FindsRoom() ([]model.Room, error) {
	var rooms []model.Room
	err := r.db.Preload("admin").Find(&rooms).Error
	return rooms, err
}

func (r *repository) GetRoom(id int) (model.Room, error) {
	var room model.Room
	err := r.db.Preload("Chat.From").Preload("Chat.Room").Preload("Admin").First(&room, id).Error
	return room, err
}

func (r *repository) UpdateRoom(room model.Room) (model.Room, error) {
	err := r.db.Save(&room).Error
	return room, err
}

func (r *repository) DeleteRoom(room model.Room) (model.Room, error) {
	err := r.db.Delete(&room).Error
	return room, err
}
