package model

type Room struct {
	ID      int        `json:"id" grom:"primary_key:auto_increment"`
	Name    string     `json:"name" gorm:"type:text"`
	Chat    []Chat     `json:"chat" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AdminID int        `json:"-"`
	Admin   UserRespon `json:"admin" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type RoomRespon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (RoomRespon) TableName() string {
	return "rooms"
}
