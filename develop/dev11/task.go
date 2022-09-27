package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	events = make(map[int]event)
)

func main() {
	startServer()
}

type middleware func(http.HandlerFunc) http.HandlerFunc

type resultJSON struct {
	Result any `json:"result"`
}

type errorJSON struct {
	ErrorMsg string `json:"error"`
}

type date time.Time

type event struct {
	Id          int    `json:"event_id"`
	Date        date   `json:"date"`
	UserId      int    `json:"user_id"`
	Description string `json:"description"`
}

func startServer() {
	http.HandleFunc("/create_event", chain(createEventHandler, method(http.MethodPost), logging()))
	http.HandleFunc("/update_event", chain(updateEventHandler, method(http.MethodPost), logging()))
	http.HandleFunc("/delete_event", chain(deleteEventHandler, method(http.MethodPost), logging()))
	http.HandleFunc("/events_for_day", chain(eventsForDayHandler, method(http.MethodGet), logging()))
	http.HandleFunc("/events_for_week", chain(eventsForWeekHandler, method(http.MethodGet), logging()))
	http.HandleFunc("/events_for_month", chain(eventsForMonthHandler, method(http.MethodGet), logging()))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

// Handlers
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	idx, err := createEvent(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w, err)
		return
	}

	sendResult(w, resultJSON{
		Result: idx,
	})
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	idx, err := updateEvent(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w, err)
		return
	}

	sendResult(w, resultJSON{
		Result: idx,
	})
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	idx, err := deleteEvent(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w, err)
		return
	}

	sendResult(w, resultJSON{
		Result: idx,
	})
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	dayQuery := r.URL.Query().Get("date")
	if len(dayQuery) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w, errors.New("you haven't specified date"))
		return
	}

	day, err := time.Parse("2006-01-02", dayQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w, errors.New("you haven't specified date"))
		return
	}

	weekEvents := getEventsInInterval(day, day.AddDate(0, 0, 1))
	sendResult(w, resultJSON{
		Result: weekEvents,
	})
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	dayQuery := r.URL.Query().Get("date")
	if len(dayQuery) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w, errors.New("you haven't specified date"))
		return
	}

	day, err := time.Parse("2006-01-02", dayQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w, errors.New("you haven't specified date"))
		return
	}

	weekEvents := getEventsInInterval(day, day.AddDate(0, 0, 7))
	sendResult(w, resultJSON{
		Result: weekEvents,
	})
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	dayQuery := r.URL.Query().Get("date")
	if len(dayQuery) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w, errors.New("you haven't specified date"))
		return
	}

	day, err := time.Parse("2006-01-02", dayQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendError(w, errors.New("you haven't specified date"))
		return
	}

	weekEvents := getEventsInInterval(day, day.AddDate(0, 1, 0))
	sendResult(w, resultJSON{
		Result: weekEvents,
	})
}

// Middlewares

// Logging logs all requests with its path and the time it took to process
func logging() middleware {

	// Create a new middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.Method, r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func method(m string) middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			var sb strings.Builder

			// Do middleware things
			if r.Method != m {
				sb.WriteString("Wrong method ")
				sb.WriteString(m)

				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(errorJSON{
					ErrorMsg: sb.String(),
				})
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

func chain(f http.HandlerFunc, middlewares ...middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Date marshall/unmarshall methods
func (d *date) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*d = date(t) //set result using the pointer
	return nil
}

func (d *date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(*d).Format("2006-01-02") + `"`), nil
}

// Utils
func sendResult(w http.ResponseWriter, data any) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		sendError(w, err)
	}
}

func createEvent(reader io.Reader) (int, error) {
	var e event

	if err := json.NewDecoder(reader).Decode(&e); err != nil {
		return -1, err
	}

	if _, ok := events[e.Id]; ok {
		return -1, errors.New("event with this id has already been created")
	}

	events[e.Id] = e

	return e.Id, nil
}

func updateEvent(reader io.Reader) (int, error) {
	var e event

	if err := json.NewDecoder(reader).Decode(&e); err != nil {
		return -1, err
	}

	if _, ok := events[e.Id]; !ok {
		return -1, errors.New(fmt.Sprintf("no event with given id [%d] have been found", e.Id))
	}

	events[e.Id] = e

	return e.Id, nil
}

func deleteEvent(reader io.Reader) (int, error) {
	var e event

	if err := json.NewDecoder(reader).Decode(&e); err != nil {
		return -1, err
	}

	if _, ok := events[e.Id]; !ok {
		return -1, errors.New(fmt.Sprintf("no event with given id [%d] have been found", e.Id))
	}

	delete(events, e.Id)

	return e.Id, nil
}

func getEventsInInterval(from time.Time, to time.Time) []event {
	var res []event

	for _, e := range events {
		if from.After(time.Time(e.Date)) && to.Before(time.Time(e.Date)) {
			res = append(res, e)
		}
	}

	return res
}

func sendError(w http.ResponseWriter, err error) {
	_ = json.NewEncoder(w).Encode(errorJSON{
		ErrorMsg: err.Error(),
	})
	log.Println(err)
}
