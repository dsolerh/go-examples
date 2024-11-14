package api

import (
	"chores_app/data"
)

func transformChore(choreData *data.ChoreDefinition) *ChoreDTO {
	return &ChoreDTO{
		ID:                    choreData.ID,
		Name:                  choreData.Name,
		Description:           choreData.Description,
		AsignedMembers:        choreData.AsignedMembers,
		ReviewExpiration:      choreData.ReviewExpiration,
		ReviewExtraExpiration: choreData.ReviewExtraExpiration,
		Priority:              choreData.Priority,
		Schedule:              transformSchedule(&choreData.Schedule),
	}
}

func transformSchedule(scheduleData *data.Schedule) *ScheduleDTO {
	return &ScheduleDTO{
		ExactDate:   scheduleData.ExactDate,
		TimesPerDay: scheduleData.TimesPerDay,
		Frequency:   scheduleData.Frequency,
		Type:        scheduleData.Type,
	}
}
