package model

type Chat struct {
	ID       int        `json:"id" gorm:"primary_key:auto_increment"`
	Message  string     `json:"message" gorm:"type:text"`
	FromUser int        `json:"-"`
	From     UserRespon `json:"form" gorm:"foreignKey:FromUser;constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	RoomID   int        `json:"-"`
	Room     RoomRespon `json:"room" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ChatResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	From    User   `json:"from"`
}

func (ChatResponse) TableName() string {
	return "chats"
}
