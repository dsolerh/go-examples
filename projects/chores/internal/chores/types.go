package chores

import "time"

type Details struct {
	Name                  string        `json:"name,omitempty"`
	Description           string        `json:"description,omitempty"`
	AsignedMembers        []string      `json:"asigned_members,omitempty"`
	ReviewExpiration      time.Duration `json:"review_expiration,omitempty"`
	ReviewExtraExpiration time.Duration `json:"review_extra_expiration,omitempty"`
	Priority              int           `json:"priority,omitempty"`
}
