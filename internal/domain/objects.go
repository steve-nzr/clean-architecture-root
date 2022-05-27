package domain

import "time"

type CreateMovie struct {
	Title      string
	ReleasedAt time.Time
}
