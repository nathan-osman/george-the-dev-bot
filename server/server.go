package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hectane/go-asyncserver"
	"github.com/nathan-osman/go-sechat"
)

// Server listens for HTTP connections and provides a basic interface for
// controlling the bot. It also provides a status endpoint for the monitoring
// system.
type Server struct {
	server *server.AsyncServer
	conn   *sechat.Conn
}

// New creates a new server.
func New(conn *sechat.Conn, addr string) (*Server, error) {
	var (
		mux = mux.NewRouter()
		s   = &Server{
			server: server.New(addr),
			conn:   conn,
		}
	)
	s.server.Handler = mux
	mux.HandleFunc("/", s.status)
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
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
