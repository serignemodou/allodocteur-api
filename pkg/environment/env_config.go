package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DatabaseHost     = "127.0.0.1"
	DatabasePort     = "3306"
	DatabaseType     = "mysql"
	DatabaseUser     = "root"
	DatabaseName     = "allodocteur"
	DatabasePassword = "password"
	ServerPort       = 8080
	ServerURL        = "http://localhost:8080"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		_ = godotenv.Load(".env")
		log.Fatal("Error loading .env file")
	}
	DatabaseHost = getStringValue("DB_HOST", DatabaseHost)
	DatabasePort = getStringValue("DB_PORT", DatabasePort)
	DatabaseType = getStringValue("DB_TYPE", DatabaseType)
	DatabaseUser = getStringValue("DB_USER", DatabaseUser)
	DatabaseName = getStringValue("DB_NAME", DatabaseName)
	DatabasePassword = getStringValue("DB_PASSWORD", DatabasePassword)
	ServerURL = getUrl("SERVER_URL", ServerURL)

	log.Println("init env config success")
}

func getStringValue(key, defaultValue string) string {
	if len(os.Getenv(key)) == 0 {
		return defaultValue
	}
	return os.Getenv(key)
}

func getUrl(key, defaultString string) string {
	result := os.Getenv(key)
	if len(result) == 0 {
		return defaultString
	}
	lastChart := result[len(result)-1:]
	if lastChart == "/" {
		result = result[:len(result)-1]
	}
	return result
}
