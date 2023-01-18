package model

type Tiket struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type:text"`
	Stok  int    `json:"Stok" gorm:"type:int"`
	Harga int    `json:"harga" gorm:"type:int"`
}
