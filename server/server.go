package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hectane/go-asyncserver"
)

// Server listens for HTTP connections and provides a basic interface for
// controlling the bot. It also provides a status endpoint for the monitoring
// system.
type Server struct {
	server *server.AsyncServer
}

// New creates a new server.
func New(addr string) (*Server, error) {
	var (
		mux = mux.NewRouter()
		s   = &Server{
			server: server.New(addr),
		}
	)
	s.server.Handler = mux
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	if err := s.server.Start(); err != nil {
		return nil, err
	}
	return s, nil
}

// Close the server.
func (s *Server) Close() {
	s.server.Stop()
}
