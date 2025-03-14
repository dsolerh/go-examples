package availability

import (
	"reservations_api/internal/availability"
	"time"

	"github.com/gofiber/fiber/v3"
)

func GetAvailability(service availability.AvailabilityService) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		slots, err := service.GetAvailability(c.Context(), time.Now(), 2)
		if err != nil {
			return nil
		}
		return c.JSON(slots)
	}
}
