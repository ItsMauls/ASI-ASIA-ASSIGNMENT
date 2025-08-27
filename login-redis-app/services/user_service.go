package services

import (
	"encoding/json"
	"errors"
	"login-system/config"
	"login-system/models"
	"login-system/utils"

	"github.com/go-redis/redis/v8"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// membuat user dengan data hardcode
func (s *UserService) CreateUser(username string) error {
	// Data hardcode
	realname := "Aberto Doni Sianturi"
	email := "adss@gmail.com"
	password := "password123"

	// Hash password dengan SHA1
	hashedPassword := utils.HashSHA1(password)

	// struct user
	user := models.User{
		RealName: realname,
		Email:    email,
		Password: hashedPassword,
	}

	// convert ke JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// simpan ke Redis
	redisKey := "login_" + username
	err = config.RedisClient.Set(config.Ctx, redisKey, userJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// memvalidasi login user
func (s *UserService) LoginUser(username, password string) (*models.User, error) {
	// Hash password yang diinput dengan SHA1
	hashedPassword := utils.HashSHA1(password)

	// ambil data user dari Redis
	redisKey := "login_" + username
	userData, err := config.RedisClient.Get(config.Ctx, redisKey).Result()

	if err == redis.Nil {
		return nil, errors.New("username tidak ditemukan")
	} else if err != nil {
		return nil, err
	}

	// parse JSON data dari Redis
	var user models.User
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return nil, err
	}

	// bandingkan password
	if user.Password != hashedPassword {
		return nil, errors.New("password salah")
	}

	return &user, nil
}
