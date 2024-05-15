package repository

import (
	calendarapi "calendar_api"
	"fmt"
	"sync"
	"time"
)

// используем мапу вместо бд
type Repository struct {
	eventsStore map[int]calendarapi.Event
	mu          sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		eventsStore: make(map[int]calendarapi.Event),
	}
}

func (r *Repository) AddEventToStore(event calendarapi.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Check if the event ID already exists in the events store
	if _, exists := r.eventsStore[event.ID]; exists {
		return fmt.Errorf("event with ID %d already exists", event.ID)
	}
	r.eventsStore[event.ID] = event
	// fmt.Println("event Added successfully")
	return nil
}

func (r *Repository) DeleteEventFromStore(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Проверяем, существует ли событие с указанным ID
	if _, exists := r.eventsStore[id]; !exists {
		return fmt.Errorf("event with id %d not found", id)
	}
	// Удаляем событие из мапы
	delete(r.eventsStore, id)
	// fmt.Println("Event with id", id, "deleted successfully")
	r.PrintMapStore()
	return nil
}
func (r *Repository) UpdateEventInStore(id int, name string, userId int, date time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	event, ok := r.eventsStore[id]
	if !ok {
		return fmt.Errorf("event with id %d not found", id)
	}
	if name != "" {
		// fmt.Println("поменяли name")
		event.Name = name

	}
	if !date.IsZero() {
		event.Date = date
	}
	if userId != 0 {
		event.UserID = userId
	}
	//обновили
	r.eventsStore[id] = event
	r.PrintMapStore()
	return nil
}
func (r *Repository) GetEventsForDayFromStore(date time.Time) ([]calendarapi.Event, error) {
	var res []calendarapi.Event
	// Обрезаем время в дате, чтобы учесть только день
	date = date.Truncate(24 * time.Hour)

	for _, event := range r.eventsStore {
		// Обрезаем время в дате события
		eventDate := event.Date.Truncate(24 * time.Hour)
		if eventDate.Equal(date) {
			res = append(res, event)
		}
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("no events found for the specified date: %s", date.Format("2006-01-02"))
	}

	return res, nil

}

func (r *Repository) GetEventsForWeekFromStore(date time.Time) ([]calendarapi.Event, error) {
	monday := date
	var res []calendarapi.Event
	// Обрезаем время в дате, чтобы учесть только день
	date = date.Truncate(24 * time.Hour)

	for i := 0; i < 7; i++ {
		// fmt.Println("зашел в цикл")
		for _, event := range r.eventsStore {
			// Обрезаем время в дате события
			eventDate := event.Date.Truncate(24 * time.Hour)
			// fmt.Println("event date", eventDate)
			if eventDate.Equal(date) {
				res = append(res, event)
			}
		}
		date = date.Add(24 * time.Hour)
		// fmt.Println("date after adding day", date)
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("no events found for the specified week starting from: %s", monday.Format("2006-01-02"))
	}

	return res, nil
}
func (r *Repository) GetEventsForMonthFromStore(firstDayOfMonth time.Time, lastDayOfMonth time.Time) ([]calendarapi.Event, error) {
	var res []calendarapi.Event

	// Проходимся по всем событиям в хранилище
	for _, event := range r.eventsStore {
		// Проверяем, находится ли дата события внутри месяца
		if event.Date.Equal(firstDayOfMonth) || event.Date.After(firstDayOfMonth) && event.Date.Before(lastDayOfMonth.AddDate(0, 0, 1)) {
			// Добавляем событие в результат
			res = append(res, event)
		}
	}

	// Если не найдено событий, возвращаем ошибку
	if len(res) == 0 {
		return nil, fmt.Errorf("no events found for the month: %s", firstDayOfMonth.Format("January 2006"))
	}

	return res, nil
}

func (r *Repository) GetStoreLen() int {
	return len(r.eventsStore)
}

func (r *Repository) PrintMapStore() {
	for key, value := range r.eventsStore {
		fmt.Printf("key is %d\n", key)
		fmt.Printf("id is %v\n", value.ID)
		fmt.Printf("name is %v\n", value.Name)
		fmt.Printf("userID is %v\n", value.UserID)
		fmt.Printf("date is %v\n", value.Date)
	}
}
