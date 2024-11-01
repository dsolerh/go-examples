package chores

import "time"

type ChoreDefinition struct {
	ID                    string
	Name                  string
	Description           string
	AsignedMembers        []string
	ReviewExpiration      time.Duration
	ReviewExtraExpiration time.Duration
	Priority              int
	Active                bool
}

type Schedule struct {
	ExactDate    time.Time
	TimesPerDay  []string
	Frequency    []int
	AmountPerDay int
	Type         ScheduleType
}

type ScheduleType uint8

type Member struct {
	ID     string
	Name   string
	Role   Role
	Active bool
}

type Role uint8

const (
	RoleAdmin Role = iota
	Reviewer
	Basic
)
