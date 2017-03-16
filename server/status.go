package server

import (
	"html/template"
	"net/http"
)

const statusHTML = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>George the Dev Bot</title>
        <style>
        body {
            font-family: sans-serif;
        }
        </style>
    </head>
    <body>
        <h1>Status Page</h1>
        <p>The bot is in the following rooms:</p>
        <ul>
            {{ range $r := .user.Rooms }}
                <li>
                    <a href="https://chat.stackexchange.com/rooms/{{ $r.ID }}">
                        {{ $r.Name }}
                    </a>
                </li>
            {{ end }}
        </ul>
    </body>
</html>
`

var statusTemplate = template.New("status")

func init() {
	if _, err := statusTemplate.Parse(statusHTML); err != nil {
		panic(err)
	}
}

// status displays status information.
func (s *Server) status(w http.ResponseWriter, r *http.Request) {
	u, err := s.conn.User(s.conn.UserID())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	statusTemplate.Execute(w, map[string]interface{}{
		"user": u,
	})
}
