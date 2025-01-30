package config

import (
	"os"
	"sync"
)

type Config struct {
	ServerPort      string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	JWTSecret       string
	MINIOEndpoint   string
	MINIOAccessKey  string
	MINIOSecretKey  string
	MINIOBucketName string
}

var (
	cfg  *Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{
			ServerPort:      getEnv("SERVER_PORT", "8080"),
			DBHost:          getEnv("DB_HOST", "localhost"),
			DBPort:          getEnv("DB_PORT", "5432"),
			DBUser:          getEnv("DB_USER", "postgres"),
			DBPassword:      getEnv("DB_PASSWORD", "postgres"),
			DBName:          getEnv("DB_NAME", "clinic"),
			JWTSecret:       getEnv("JWT_SECRET", "default-secret"),
			MINIOEndpoint:   getEnv("MINIO_ENDPOINT", "localhost:9000"),
			MINIOAccessKey:  getEnv("MINIO_ACCESS_KEY", "minioadmin"),
			MINIOSecretKey:  getEnv("MINIO_SECRET_KEY", "minioadmin"),
			MINIOBucketName: getEnv("MINIO_BUCKET_NAME", "patients"),
		}
	})
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
