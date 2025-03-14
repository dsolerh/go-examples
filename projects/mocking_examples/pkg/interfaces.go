package pkg

import "time"

type Interface interface {
	GetValue(int, string, time.Duration) (string, error)
}
