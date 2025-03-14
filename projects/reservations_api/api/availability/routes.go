package availability

import (
	"reservations_api/internal/availability"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router) {
	service := availability.NewService()

	router.Get("", GetAvailability(service))
}
