package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/siddheshRajendraNimbalkar/Go-Ecommerce/BACKEND/internal/helper"
	"gorm.io/gorm"
)

type RestRoutes struct {
	App  *fiber.App
	DB   *gorm.DB
	Auth helper.Auth
}
