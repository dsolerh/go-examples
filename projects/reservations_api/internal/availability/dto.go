package availability

import "time"

type AvailabilitySlot struct {
	ScheduleId int       `json:"schedule_id"`
	TableId    int       `json:"table_id"`
	Capacity   int       `json:"capacity"`
	Date       time.Time `json:"date"`
}
