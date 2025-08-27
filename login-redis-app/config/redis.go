package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // no password
		DB:       0,                // default DB
	})

	// test koneksi Redis
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Gagal koneksi ke Redis:", err)
	}
	fmt.Println("Berhasil terhubung ke Redis!")
}
