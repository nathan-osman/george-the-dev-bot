package registry

import (
	"github.com/nathan-osman/go-sechat"
)

// Command is run once for each event.
type Command func(*sechat.Conn, *sechat.Event)

// commands is used to maintain a list of commands.
var commands []Command

// Register adds a command to the list of commands.
func Register(cmd Command) {
	commands = append(commands, cmd)
}

// Execute runs the event through each of the commands.
func Execute(conn *sechat.Conn, event *sechat.Event) {
	for _, cmd := range commands {
		cmd(conn, event)
	}
}
