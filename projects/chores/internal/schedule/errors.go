package schedule

import "errors"

var (
	ErrInvalidScheduleType = errors.New("invalid schedule type")
	ErrInvalidHourOrMoment = errors.New("invalid hour or moment")
	ErrInvalidFrequency    = errors.New("invalid frequency")
	ErrInvalidExactDate    = errors.New("invalid exact date")
)
