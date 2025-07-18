package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	JWTSecret      string
	JWTExpireHours int
	RedisHost      string
	RedisPort      string
	RedisPassword  string
	RedisDB        int
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	expire, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS"))

	return &Config{
		AppPort:        os.Getenv("APP_PORT"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		JWTExpireHours: expire,
		RedisHost:      os.Getenv("REDIS_HOST"),
		RedisPort:      os.Getenv("REDIS_PORT"),
		RedisPassword:  os.Getenv("REDIS_PASSWORD"),
		RedisDB:        func() int { v, _ := strconv.Atoi(os.Getenv("REDIS_DB")); return v }(),
	}
}