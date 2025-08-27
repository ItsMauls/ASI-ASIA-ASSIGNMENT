package models

// struct untuk data user di Redis
type User struct {
	RealName string `json:"realname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// struct untuk request login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// struct untuk response API
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
