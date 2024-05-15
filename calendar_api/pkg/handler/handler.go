package handler

import (
	"calendar_api/pkg/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.createEventHandler)
	mux.HandleFunc("/update_event", h.updateEventHandler)
	mux.HandleFunc("/delete_event", h.deleteEventHandler)
	mux.HandleFunc("/events_for_day", h.eventsForDayHandler)
	mux.HandleFunc("/events_for_week", h.eventsForWeekHandler)
	mux.HandleFunc("/events_for_month", h.eventsForMonthHandler)
	// Добавляем middleware для логирования запросов
	handlerWithLogging := h.loggingMiddleware(mux)

	return handlerWithLogging

}
func (h *Handler) createEventHandler(w http.ResponseWriter, r *http.Request) {
	//парсим тело запроса
	name, userId, date, err := h.parseDataForPost(r)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	//вызываем соотвествующую функцию сервиса
	err = h.services.AddEvent(name, userId, date)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	sendJSONResult(w, "Event Added Successfully")
}

func (h *Handler) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseId(r)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	//значит получаем id
	//также забираем то что нужно изменить: name, date, or iserID
	//парсим айди
	name, userId, date, err := h.parseDataForPost(r)
	// fmt.Println("спарсенные данные", name, userId, date)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.services.UpdateEvent(id, name, userId, date)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	sendJSONResult(w, "event updated successuflly")
}

func (h *Handler) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseId(r)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.services.DeleteEvent(id)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	sendJSONResult(w, "event deleted successuflly")
}

func (h *Handler) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	date, err := h.parseDateFromQuery(r)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.services.GetEventsForDay(date)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	sendJSONResult(w, res)
}

func (h *Handler) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	date, err := h.parseDateFromQuery(r)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	//пока дата не станет понедельником текущей недели, вычитаем по дню
	for date.Weekday() != time.Monday {
		date = date.Add(-24 * time.Hour)
	}

	// fmt.Println(date.Weekday())
	//теперь передаем дату в функцию соответствующую из сервиса в следующий слой
	eventsForWeek, err := h.services.GetEventsForWeek(date)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	sendJSONResult(w, eventsForWeek)
}

func (h *Handler) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	date, err := h.parseDateFromQuery(r)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		sendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	firstDayOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	fmt.Println(firstDayOfMonth)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)
	fmt.Println(lastDayOfMonth)
	res, err := h.services.GetEventsForMonth(firstDayOfMonth, lastDayOfMonth)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	fmt.Println(res)
	sendJSONResult(w, res)
}

// HELPERS
func (h *Handler) parseDataForPost(r *http.Request) (name string, userID int, date time.Time, err error) {
	var userIdstring string
	var dateString string

	err = r.ParseForm()
	if err != nil {
		return "", 0, time.Time{}, errors.New("failed to parse data")
	}
	//спарсили юзерайди
	userIdstring = r.Form.Get("user_id")
	if userIdstring != "" {
		userID, err = strconv.Atoi(userIdstring)
		if err != nil {
			return "", 0, time.Time{}, errors.New("failed to convert userId to string")
		}
	}
	//спарсили дату
	dateString = r.Form.Get("date")
	if dateString != "" {
		date, err = time.Parse("2006-01-02", dateString)
		if err != nil {
			return "", 0, time.Time{}, errors.New("failed to convert date to time format")
		}
	}
	//спарсили имя
	name = r.Form.Get("name")
	return name, userID, date, nil
}
func (h *Handler) parseId(r *http.Request) (int, error) {
	var id int
	var err error
	idString := r.FormValue("id")
	if idString != "" {
		id, err = strconv.Atoi(idString)
		if err != nil {
			return 0, fmt.Errorf("failed to convert id to int")
		}
	}
	return id, nil
}

func (h *Handler) parseDateFromQuery(r *http.Request) (time.Time, error) {
	// Получаем значения параметров из строки запроса
	queryValues := r.URL.Query()

	// Получаем значение параметра "date"
	dateString := queryValues.Get("date")

	// Проверяем, не является ли значение параметра "date" пустым
	if dateString == "" {
		return time.Time{}, errors.New("date parameter is missing")
	}

	// Преобразуем значение параметра "date" в формат time.Time
	parsedDate, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return time.Time{}, errors.New("failed to parse date parameter")
	}

	return parsedDate, nil
}

// /json respond senders

func sendJSONResult(w http.ResponseWriter, result interface{}) {
	response := map[string]interface{}{"result": result}
	sendJSONResponse(w, http.StatusOK, response)
}
func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func sendJSONError(w http.ResponseWriter, errorMsg string, statusCode int) {
	response := map[string]string{"error": errorMsg}
	sendJSONResponse(w, statusCode, response)
}
