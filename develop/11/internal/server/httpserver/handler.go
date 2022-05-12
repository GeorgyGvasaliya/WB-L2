package httpserver

import (
	"11/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("expect method POST at /create_event, got %v", r.Method), true)
		return
	}
	event := &model.Event{}
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, err.Error(), true)
		return
	}

	if err := s.Calendar.CreateEvent(r.Context(), *event); err != nil {
		writeJSONResponse(w, http.StatusServiceUnavailable, err.Error(), true)
		return
	}
	writeJSONResponse(w, http.StatusCreated, fmt.Sprintf("event created"), false)
}
func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("expect method POST at /delete_event, got %v", r.Method), true)
		return
	}
	event := &model.Event{}
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, err.Error(), true)
		return
	}
	if event.ID == 0 {
		writeJSONResponse(w, http.StatusInternalServerError, "invalid id event", true)
		return
	}
	if err := s.Calendar.DeleteEvent(r.Context(), event.ID); err != nil {
		writeJSONResponse(w, http.StatusServiceUnavailable, err.Error(), true)
		return
	}
	writeJSONResponse(w, http.StatusOK, fmt.Sprintf("event deleted"), false)
}
func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("expect method POST at /update_event, got %v", r.Method), true)
		return
	}

	event := &model.Event{}
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, err.Error(), true)
		return
	}

	if event.ID == 0 {
		writeJSONResponse(w, http.StatusInternalServerError, "invalid id event", true)
		return
	}

	if err := s.Calendar.UpdateEvent(r.Context(), *event); err != nil {
		writeJSONResponse(w, http.StatusServiceUnavailable, err.Error(), true)
		return
	}
	writeJSONResponse(w, http.StatusOK, fmt.Sprintf("event updated"), false)
}

func (s *Server) eventsToday(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("expect method GET at /events_for_day, got %v", r.Method), true)
		return
	}
	events, err := s.Calendar.GetEvent(r.Context(), "day")
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, err.Error(), true)
		return
	}
	writeJSONEvents(w, http.StatusOK, events)

}
func (s *Server) eventsThisWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("expect method GET at /events_for_week, got %v", r.Method), true)
		return
	}
	events, err := s.Calendar.GetEvent(r.Context(), "week")
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, err.Error(), true)
		return
	}

	writeJSONEvents(w, http.StatusOK, events)
}
func (s *Server) eventsThisMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("expect method GET at /events_for_month, got %v", r.Method), true)
		return
	}
	events, err := s.Calendar.GetEvent(r.Context(), "month")
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, err.Error(), true)
		return
	}
	writeJSONEvents(w, http.StatusOK, events)
}
