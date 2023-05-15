package domain

import "time"

type Ad struct {
	Id, Title, Description string
	Price                  uint32
	Date                   time.Time
}
