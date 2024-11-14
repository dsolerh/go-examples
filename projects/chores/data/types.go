package data

import "time"

type ChoreDetails struct {
	Name                  string
	Description           string
	AsignedMembers        []string
	ReviewExpiration      time.Duration
	ReviewExtraExpiration time.Duration
	Priority              int
}

// table: chore_definitions
type ChoreDefinition struct {
	ID string
	ChoreDetails
	Schedule Schedule
	Active   bool
}

type Schedule struct {
	ExactDate   time.Time
	TimesPerDay []string
	Frequency   []uint8
	Type        ScheduleType
}

type ScheduleType = uint8

type MemberInfo struct {
	Name string
	Role Role
}

// table: members
type Member struct {
	ID string
	MemberInfo
	Active bool
}

// table: notifications

type Role uint8

const (
	RoleAdmin Role = iota
	Reviewer
	Basic
)
