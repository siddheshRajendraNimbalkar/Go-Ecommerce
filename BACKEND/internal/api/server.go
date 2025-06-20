package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/configs"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/api/rest"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/api/rest/routes"
)

type Server struct {
	fiber  *fiber.App
	config configs.Config
}

func NewServer(config configs.Config) *Server {

	fiberApp := fiber.New()

	r := rest.RestRoutes{
		App: fiberApp,
	}

	SetupRoutes(&r)

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

func SetupRoutes(r *rest.RestRoutes) {
	routes.SetupUserRoutes(r)
}
