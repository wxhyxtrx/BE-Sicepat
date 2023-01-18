package roomchatdto

type RequestRoomChat struct {
	Nameroom string `json:"nameroom" validate:"required"`
}
