package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	SRV         ServerConfig
	DB          DbConfig
	BYCRYPTCOST string
}

type ServerConfig struct {
	Port string
}

type DbConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
	SslMode  string
}

// Возвращает DSN как строку
func (dbc DbConfig) GetDsn() string {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", dbc.Host, dbc.User, dbc.Password, dbc.DbName, dbc.Port, dbc.SslMode)
	return dsn
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла")
	}

	return &Config{SRV: ServerConfig{Port: getEnv("SERVER_PORT", "8080")},
		DB: DbConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			User:     getEnv("POSTGRES_USER", "user"),
			Password: getEnv("POSTGRES_PASSWORD", "password"),
			DbName:   getEnv("POSTGRES_DB", "education"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			SslMode:  getEnv("SSL_MODE", "disable"),
		},
		BYCRYPTCOST: getEnv("BCRYPT_COST", "10"),
	}
}

func getEnv(key, defaultValue string) string {
	s := os.Getenv(key)
	if s == "" {
		log.Println("Применилось значение по умолчанию для ключа", key)
		return defaultValue
	}
	return s
}
