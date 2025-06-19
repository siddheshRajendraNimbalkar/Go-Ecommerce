package main

import (
	"fmt"
	"log"

	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/configs"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/api"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	server := api.NewServer(*config)
	fmt.Println("Server is running on port", config.Port)

	api.StartServer(server)
}
