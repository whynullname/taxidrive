package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
	router chi.Router
}

func NewServer(handlers *HandlersProvider, endPointAddress string) *Server {
	router := chi.NewRouter()
	handlers.ConnectHandlers(router)
	server := &http.Server{
		Addr:    endPointAddress,
		Handler: router,
	}

	return &Server{
		router: router,
		server: server,
	}
}

// TODO: Gracefull shutdown
func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}
