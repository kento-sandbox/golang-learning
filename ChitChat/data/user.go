package data

import "time"

type Thread struct {
	id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}
