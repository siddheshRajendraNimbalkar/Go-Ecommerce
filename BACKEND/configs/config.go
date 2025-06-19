package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() (config *Config, err error) {

	godotenv.Load()
	if strings.ToLower(os.Getenv("APP_ENV")) == "dev" {
		fmt.Println("Running in dev mode")
	}

	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) < 1 {
		return nil, fmt.Errorf("port address not found")
	}

	return &Config{
		Port: httpPort,
	}, nil

}
