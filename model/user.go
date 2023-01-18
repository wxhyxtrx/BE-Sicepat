package model

import "time"

type User struct {
	ID       int       `json:"id" gorm:"primary_key:auto_increment"`
	Username string    `json:"username" gorm:"type:varchar(255)"`
	Email    string    `json:"email" gorm:"type:varchar(255)"`
	Password string    `json:"password" gorm:"type:varchar(255)"`
	Status   string    `json:"status" gorm:"type:text"`
	TimeOFF  time.Time `json:"time_off"`
}

type UserRespon struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (UserRespon) TableName() string {
	return "users"
}
