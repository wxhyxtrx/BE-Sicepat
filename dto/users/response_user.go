package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Status   string `json:"status" from:"status"`
}
