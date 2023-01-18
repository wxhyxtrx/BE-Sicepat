package tiketdto

type RequestTiket struct {
	Name  string `json:"name" validate:"required"`
	Stok  int    `json:"stok" validate:"required"`
	Harga int    `json:"harga" validate:"required"`
}
