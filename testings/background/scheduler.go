package main

import "time"

var _ Scheduler = (*scheduler)(nil)

type SchedulerOptions struct {
	// if specified it will be used to create a one of background task
	// to be run at the specified timestamp, if the timestamp is before
	// the current time it will not schedule the task
	ExecuteAt time.Time
	// if specified it will be used to create a recurring background task
	// that will run at the interval defined by the param
	Duration time.Duration
	// if specified it will be used to stop the task execution after the provided
	// amount of time. It will not work with ExecuteAt and if both are specified
	// this value will be ignored
	ExpiresIn time.Duration
	// if set to true it will trigger a response from the channel to indicate that the
	// background task should be immediately start running. It wont work with ExecuteAt.
	ImmediateStart bool
}

func NewScheduleBuilder(options SchedulerOptions) Scheduler {
	return &scheduler{
		duration:  options.Duration,
		expiresIn: options.ExpiresIn,
		executeAt: options.ExecuteAt,
		immediate: options.ImmediateStart,
	}
}

type scheduler struct {
	executeAt time.Time
	duration  time.Duration
	expiresIn time.Duration
	immediate bool
}

// New implements Scheduler.
func (s *scheduler) New() schedule {
	if !s.executeAt.IsZero() {
		if s.executeAt.Before(time.Now()) {
			return neverSchedule{}
		}
		return newOneOffSchedule(s.executeAt)
	}
	return newRecurringSchedule(s.duration, s.expiresIn, s.immediate)
}
