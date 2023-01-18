package chatdto

type RequestChat struct {
	Message string `json:"message"`
	RoomID  int    `json:"roomid"`
}
