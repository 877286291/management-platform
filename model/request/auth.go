package request

type LoginRequest struct {
	LoginName     string `json:"login_name"`
	LoginPassword string `json:"login_password"`
}
