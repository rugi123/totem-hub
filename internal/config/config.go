package config

import (
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Version     string
	LoggerLevel string
	JWTKey      string
}

type Postgres struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	SSLMode    string
}

type Config struct {
	App      App
	Postgres Postgres
}

func InitConfig(path string) (Config, error) {
	err := godotenv.Load(path)
	return Config{
		App: App{
			Version:     os.Getenv("APP_VERSION"),
			LoggerLevel: os.Getenv("LOGGER_LEVEL"),
			JWTKey:      os.Getenv("JWT_SECRET"),
		},
		Postgres: Postgres{
			DBUser:     os.Getenv("POSTGRES_USER"),
			DBPassword: os.Getenv("POSTGRES_PASSWORD"),
			DBHost:     os.Getenv("POSTGRES_HOST"),
			DBPort:     os.Getenv("POSTGRES_PORT"),
			DBName:     os.Getenv("POSTGRES_DB_NAME"),
			SSLMode:    os.Getenv("POSTGRES_SSL_MODE"),
		},
	}, err
}
