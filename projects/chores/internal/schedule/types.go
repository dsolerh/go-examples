package schedule

import (
	"time"
)

type Schedule struct {
	ExactDate   time.Time      `json:"exact_date,omitempty"`
	TimesPerDay []HourOrMoment `json:"times_per_day,omitempty"`
	Frequency   []day          `json:"frequency,omitempty"`
	Type        Type           `json:"type,omitempty"`
}

// can represent a day in a week (1-7)
// of in a month (1-30), will not include 31 and will consider all
// months to have 30 days
type day = uint8

type HourOrMoment = string

const (
	Morning HourOrMoment = "morning"
	MidDay  HourOrMoment = "mid-day"
	Evening HourOrMoment = "evening"
)

type Type = uint8

const (
	Daily Type = iota
	Monthly
	ExactDate
)
