package handlers

import (
	"login-system/models"
	"login-system/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userService: services.NewUserService(),
	}
}

// menangani pembuatan user dengan data hardcode
func (h *AuthHandler) CreateUser(c *fiber.Ctx) error {
	// Ambil username dari query parameter
	username := c.Query("username")
	if username == "" {
		return c.Status(400).JSON(models.Response{
			Success: false,
			Message: "Username harus diisi di query parameter",
		})
	}

	// Buat user menggunakan service
	err := h.userService.CreateUser(username)
	if err != nil {
		return c.Status(500).JSON(models.Response{
			Success: false,
			Message: "Error membuat user: " + err.Error(),
		})
	}

	// Response berhasil
	responseData := map[string]any{
		"username": username,
		"realname": "Aberto Doni Sianturi",
		"email":    "adss@gmail.com",
	}

	return c.JSON(models.Response{
		Success: true,
		Message: "User berhasil dibuat",
		Data:    responseData,
	})
}

// menangani proses login
func (h *AuthHandler) ProcessLogin(c *fiber.Ctx) error {
	// Parse JSON request
	var loginReq models.LoginRequest
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(400).JSON(models.Response{
			Success: false,
			Message: "Format JSON tidak valid",
		})
	}

	// validasi input kosong
	if loginReq.Username == "" || loginReq.Password == "" {
		return c.Status(400).JSON(models.Response{
			Success: false,
			Message: "Username dan password harus diisi",
		})
	}

	// proses login menggunakan service
	user, err := h.userService.LoginUser(loginReq.Username, loginReq.Password)
	if err != nil {
		status := 401
		if err.Error() == "username tidak ditemukan" {
			status = 401
		} else if err.Error() == "password salah" {
			status = 401
		} else {
			status = 500
		}

		return c.Status(status).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	// login berhasil - jangan return password
	responseData := map[string]any{
		"username": loginReq.Username,
		"realname": user.RealName,
		"email":    user.Email,
	}

	return c.JSON(models.Response{
		Success: true,
		Message: "Login berhasil",
		Data:    responseData,
	})
}
