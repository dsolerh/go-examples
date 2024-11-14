package api

import (
	"chores_app/internal/chores"
	"chores_app/internal/schedule"
	"time"
)

type CreateChoreDTO struct {
	Details  *chores.Details    `json:"details,omitempty"`
	Schedule *schedule.Schedule `json:"schedule,omitempty"`
}

type ChoreDTO struct {
	ID                    string        `json:"id,omitempty"`
	Name                  string        `json:"name,omitempty"`
	Description           string        `json:"description,omitempty"`
	AsignedMembers        []string      `json:"asigned_members,omitempty"`
	ReviewExpiration      time.Duration `json:"review_expiration,omitempty"`
	ReviewExtraExpiration time.Duration `json:"review_extra_expiration,omitempty"`
	Priority              int           `json:"priority,omitempty"`
	Schedule              *ScheduleDTO  `json:"schedule,omitempty"`
}

type ScheduleDTO struct {
	ExactDate   time.Time `json:"exact_date,omitempty"`
	TimesPerDay []string  `json:"times_per_day,omitempty"`
	Frequency   []uint8   `json:"frequency,omitempty"`
	Type        uint8     `json:"type,omitempty"`
}

type ErrorDTO struct {
	Error   string `json:"error,omitempty"`
	Details string `json:"details,omitempty"`
}
