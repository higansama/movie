package reqres

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
