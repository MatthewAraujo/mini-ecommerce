package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Redis struct {
	Port     string
	Address  string
	Password string
	Database int
}

type JWT struct {
	JWTExpirationInSeconds int64
	JWTSecret              string
}

type API struct {
	PublicHost string
	Port       string
}

type Postgres struct {
	URL string
}

type Config struct {
	JWT      JWT
	API      API
	Redis    Redis
	Postgres Postgres
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		API: API{
			PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
			Port:       getEnv("PORT", "8080"),
		},
		JWT: JWT{
			JWTSecret:              getEnv("JWT_SECRET", "not-that-secret"),
			JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		},
		Redis: Redis{
			Port:     getEnv("REDIS_PORT", "6379"),
			Address:  getEnv("REDIS_ADDRESS", "localhost"),
			Password: getEnv("REDIS_PASSWORD", ""),
			Database: int(getEnvAsInt("REDIS_DATABASE", 0)),
		},
		Postgres: Postgres{
			URL: getEnv("POSTGRES_URL", "postgresql://docker:docker@pg:5432/ecommerce"),
		},
	}

}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
