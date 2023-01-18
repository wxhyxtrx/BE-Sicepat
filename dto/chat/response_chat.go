package chatdto

type ResponseChat struct {
	Message string `json:"message"`
	User    string `json:"user"`
	Room    string `json:"room"`
}
