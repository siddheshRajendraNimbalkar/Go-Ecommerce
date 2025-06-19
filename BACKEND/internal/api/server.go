package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/configs"
)

type Server struct {
	fiber  *fiber.App
	config configs.Config
}

func NewServer(config configs.Config) *Server {

	fiberApp := fiber.New()

	return &Server{
		fiber:  fiberApp,
		config: config,
	}
}

func StartServer(s *Server) {
	err := s.fiber.Listen(s.config.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
