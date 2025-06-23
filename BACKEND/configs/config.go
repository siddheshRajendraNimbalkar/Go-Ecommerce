package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Dsn       string
	AppSecret string
}

func LoadConfig() (config *Config, err error) {

	if strings.ToLower(os.Getenv("APP_ENV")) == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) < 1 {
		return nil, fmt.Errorf("port address not found")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return nil, fmt.Errorf("dsn not found")
	}

	AppSecret := os.Getenv("APPSECRET")
	if len(AppSecret) < 1 {
		return nil, fmt.Errorf("app secret not found")
	}

	return &Config{
		Port:      httpPort,
		Dsn:       dsn,
		AppSecret: AppSecret,
	}, nil

}
