package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVal(key, fallback string) string {
	godotenv.Load()

	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return fallback
}
