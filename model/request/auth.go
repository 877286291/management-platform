package request

type LoginRequest struct {
	LoginName     string `json:"login_name" binding:"required"`
	LoginPassword string `json:"login_password" binding:"required"`
}
