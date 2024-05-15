package calendarapi

import "time"

type Event struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
}
