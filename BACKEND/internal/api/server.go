package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/configs"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/api/rest"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/api/rest/routes"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/domain"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	fiber  *fiber.App
	config configs.Config
}

func NewServer(config configs.Config) *Server {

	fiberApp := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Printf("Connected to database successfully")

	// Migrate the schema
	db.AutoMigrate(&domain.User{})

	auth := helper.NewAuth(config.AppSecret)

	r := rest.RestRoutes{
		App:  fiberApp,
		DB:   db,
		Auth: auth,
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
