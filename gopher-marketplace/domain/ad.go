package domain

import (
	"errors"
	"time"
)

type Ad struct {
	Id, Title, Description string
	Price                  uint32
	Date                   time.Time
}

func NewAd(id string, title string, description string, price uint32) (error, *Ad) {
	if len(description) > 50 {
		return errors.New("the ad description cannot be longer than 50 characters"), nil
	}
	return nil, &Ad{
		Id:          id,
		Title:       title,
		Description: description,
		Price:       price,
		Date:        time.Now(),
	}
}
