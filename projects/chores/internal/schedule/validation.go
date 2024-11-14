package schedule

import (
	"chores_app/pkg/common"
	"time"
)

var freqlimit = [...]int{7, 30}

func (s *Schedule) Validate() error {
	switch s.Type {
	case Daily, Monthly:
		if common.Some(s.TimesPerDay, invalidHourOrMoment) {
			return ErrInvalidHourOrMoment
		}

		if len(s.Frequency) > freqlimit[s.Type] {
			return ErrInvalidFrequency
		}
		if common.Some(s.Frequency, invalidDay(s.Type)) {
			return ErrInvalidFrequency
		}
		return nil

	case ExactDate:
		if s.ExactDate.Before(time.Now()) {
			return ErrInvalidExactDate
		}
		return nil

	default:
		return ErrInvalidScheduleType
	}
}

func invalidHourOrMoment(t HourOrMoment) bool {
	switch t {
	case Morning, MidDay, Evening:
		return false
	default:
		return validhour(t)
	}
}

// validhour returns true if a string contains a valid hour
// by the format hh:mm
func validhour(s HourOrMoment) bool {
	if len(s) != 5 {
		return false
	}
	return common.IsDigit(s[0]) &&
		common.IsDigit(s[1]) &&
		s[2] == ':' &&
		common.IsDigit(s[3]) &&
		common.IsDigit(s[4])
}

func invalidDay(freq Type) func(day) bool {
	switch freq {
	case Daily:
		return func(d day) bool { return d < 1 || d > 7 }
	case Monthly:
		return func(d day) bool { return d < 1 || d > 30 }
	case ExactDate:
		panic("can only validate monthly and daily")
	default:
		panic("can only validate monthly and daily")
	}
}
