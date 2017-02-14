package registry

import (
	"github.com/nathan-osman/go-sechat"
)

// Command is run once for each event.
type Command func(*sechat.Conn, *sechat.Event) bool

// Commands are currently split into two groups: one for regular commands and
// one for "catch all" commands that will match when no other ones do (such as
// the Wikipedia ones)
const (
	RegularCommand = iota
	ReferenceCommand
)

var commands = make(map[int][]Command)

// Register adds a command to the list of commands.
func Register(cmd Command, commandType int) {
	m, _ := commands[commandType]
	commands[commandType] = append(m, cmd)
}

// Execute runs the event through each of the commands.
func Execute(conn *sechat.Conn, event *sechat.Event) {
	for _, cmd := range commands[RegularCommand] {
		if cmd(conn, event) {
			return
		}
	}
	for _, cmd := range commands[ReferenceCommand] {
		if cmd(conn, event) {
			return
		}
	}
}
