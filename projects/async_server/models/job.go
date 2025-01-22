package models

import (
	"time"

	"github.com/google/uuid"
)

type MsgType int

const (
	LogType MsgType = iota
	CallBackType
	MailType
)

func (m MsgType) String() string {
	switch m {
	case LogType:
		return "Log"
	case CallBackType:
		return "Callback"
	case MailType:
		return "Mail"
	default:
		return "Invalid"
	}
}

// Job represents UUID of a Job
type Job[T Log | CallBack | Mail] struct {
	ID        uuid.UUID `json:"uuid"`
	Type      MsgType   `json:"type"`
	ExtraData T         `json:"extra_data"`
}

// Worker-A data
type Log struct {
	ClientTime time.Time `json:"client_time"`
}

// CallBack data
type CallBack struct {
	CallBackURL string `json:"callback_url"`
}

// Mail data
type Mail struct {
	EmailAddress string `json:"email_address"`
}
