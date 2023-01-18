package model

import "time"

type Order struct {
	ID      int        `json:"id" gorm:"primary_key:auto_increment"`
	TiketID int        `json:"-"`
	Tiket   Tiket      `json:"tiket" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Qty     int        `json:"qty" gorm:"type:int"`
	Total   int        `json:"total" gorm:"type:int"`
	Tanggal time.Time  `json:"tanggal"`
	UserID  int        `json:"-"`
	User    UserRespon `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
