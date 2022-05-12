package httpserver

import (
	"11/internal/application"
	"11/internal/config"
	"context"
	"net"
	"net/http"
)

type Server struct {
	Calendar   *application.Calendar
	httpServer *http.Server
}

func New(cfg config.HTTP, App *application.Calendar) *Server {
	server := &Server{
		httpServer: &http.Server{
			Addr: net.JoinHostPort("", cfg.Port),
		},
		Calendar: App,
	}

	return server
}
func (s *Server) initRouts() http.Handler {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/create_event", s.createEvent)
	serveMux.HandleFunc("/update_event", s.updateEvent)
	serveMux.HandleFunc("/delete_event", s.deleteEvent)
	serveMux.HandleFunc("/events_for_day", s.eventsToday)
	serveMux.HandleFunc("/events_for_week", s.eventsThisWeek)
	serveMux.HandleFunc("/events_for_month", s.eventsThisMonth)
	handler := Logging(serveMux)
	return handler
}
func (s *Server) Start() error {
	s.httpServer.Handler = s.initRouts()
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
