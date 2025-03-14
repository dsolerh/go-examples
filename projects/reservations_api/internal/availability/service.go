package availability

import (
	"context"
	"time"
)

var _ AvailabilityService = (*availabilityService)(nil)

func NewService() AvailabilityService {
	return &availabilityService{}
}

type availabilityService struct {
}

// GetAvailability implements AvailabilityService.
func (a *availabilityService) GetAvailability(ctx context.Context, day time.Time, capacity int) ([]AvailabilitySlot, error) {
	return []AvailabilitySlot{
		{
			ScheduleId: 0,
			TableId:    0,
			Capacity:   capacity,
			Date:       time.Time{},
		},
		{
			ScheduleId: 0,
			TableId:    0,
			Capacity:   capacity,
			Date:       time.Time{},
		},
		{
			ScheduleId: 0,
			TableId:    0,
			Capacity:   capacity,
			Date:       time.Time{},
		},
	}, nil
}
