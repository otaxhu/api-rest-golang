package settings

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // DELETE THIS ON PRODUCTION, DO NOT USE .env FILES
)

type Server struct {
	Port      uint
	Framework string
}

type Database struct {
	User     string
	Password string
	Host     string
	Port     uint
	Driver   string
	Name     string
}

func NewServer() (*Server, error) {
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil || serverPort < 0 {
		return nil, fmt.Errorf("NewServer(): SERVER_PORT does not exists in the .env file or is not a positive integer")
	}
	serverFramework, exists := os.LookupEnv("SERVER_FRAMEWORK")
	if !exists {
		return nil, fmt.Errorf("NewServer(): SERVER_FRAMEWORK does not exists in the .env file")
	}
	return &Server{
		Port:      uint(serverPort),
		Framework: serverFramework,
	}, nil
}

func NewDatabase() (*Database, error) {
	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		return nil, fmt.Errorf("NewDatabase(): DB_USER does not exists in the .env file")
	}
	password := os.Getenv("DB_PASSWORD")
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		return nil, fmt.Errorf("NewDatabase(): DB_HOST does not exists in the .env file")
	}
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil || dbPort < 0 {
		return nil, fmt.Errorf("NewDatabase(): DB_PORT does not exists in the .env file or is not a positive integer")
	}
	driver, exists := os.LookupEnv("DB_DRIVER")
	if !exists {
		return nil, fmt.Errorf("NewDatabase(): DB_DRIVER does not exists in the .env file")
	}
	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return nil, fmt.Errorf("NewDatabase(): DB_NAME does not exists in the .env file")
	}
	return &Database{
		User:     user,
		Password: password,
		Host:     host,
		Port:     uint(dbPort),
		Driver:   driver,
		Name:     dbName,
	}, nil
}
