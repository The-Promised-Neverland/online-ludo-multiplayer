package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port     string
	Postgres PostgresConfig
	Redis    RedisConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	PoolSize int
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
	PoolSize int
}

func getEnv(key string, fallbackValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallbackValue
}

func getEnvAsInt(key string, fallbackValue int) int { 
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal 
		}
	}
	return fallbackValue
}

var AppConfigValues *AppConfig

func LoadConfig() {
	if err := godotenv.Load(); err != nil{
		log.Fatalf("Error loading .env file %v",err)
		return
	}

 	AppConfigValues = &AppConfig{
		Port: getEnv("APP_PORT", "8080"),
		Postgres: PostgresConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", "password"),
			DBName:   getEnv("POSTGRES_DB", "ludo"),
			PoolSize: getEnvAsInt("POSTGRES_POOL_SIZE", 5),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
			PoolSize: getEnvAsInt("REDIS_POOL_SIZE", 20),
		},
	}
}
