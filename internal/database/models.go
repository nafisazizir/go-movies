// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package database

import (
	"time"
)

type Movie struct {
	ID          int32
	Title       string
	Description string
	ReleaseDate time.Time
	PosterUrl   string
	AgeRating   string
	TicketPrice int32
}