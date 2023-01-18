package model

type Setting struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Teks string `json:"teks" gorm:"type:text"`
}
