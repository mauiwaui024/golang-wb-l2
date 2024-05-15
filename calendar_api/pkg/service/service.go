package service

import (
	calendarapi "calendar_api"
	"calendar_api/pkg/repository"
	"fmt"
	"time"
)

type Service struct {
	// Можно добавить бд зависисмость
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) AddEvent(name string, userID int, date time.Time) error {
	eventId := s.Repo.GetStoreLen()
	newEvent := calendarapi.Event{ID: eventId, Name: name, Date: date, UserID: userID}
	// return nil
	return s.Repo.AddEventToStore(newEvent)
}

//	func (s *Service) GetEventByID(id int) (event calendarapi.Event, ok bool) {
//		event, ok = s.Repo.GetEventByIDFromStore(id)
//		return event, ok
//	}
func (s *Service) UpdateEvent(id int, name string, userId int, date time.Time) error {
	fmt.Println("name in services")
	return s.Repo.UpdateEventInStore(id, name, userId, date)
}
func (s *Service) DeleteEvent(id int) error {
	return s.Repo.DeleteEventFromStore(id)
}

func (s *Service) GetEventsForDay(date time.Time) ([]calendarapi.Event, error) {
	return s.Repo.GetEventsForDayFromStore(date)
}

func (s *Service) GetEventsForWeek(date time.Time) ([]calendarapi.Event, error) {
	return s.Repo.GetEventsForWeekFromStore(date)
}

func (s *Service) GetEventsForMonth(firstDayOfMonth time.Time, lastDayOfMonth time.Time) ([]calendarapi.Event, error) {
	return s.Repo.GetEventsForMonthFromStore(firstDayOfMonth, lastDayOfMonth)
}
