package availability

import (
	"context"
	"time"
)

type AvailabilityService interface {
	GetAvailability(ctx context.Context, day time.Time, capacity int) ([]AvailabilitySlot, error)
}
