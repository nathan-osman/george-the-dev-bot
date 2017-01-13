package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/george-the-dev-bot/registry"
	"github.com/nathan-osman/george-the-dev-bot/server"
	"github.com/nathan-osman/go-sechat"

	_ "github.com/nathan-osman/george-the-dev-bot/net"
	_ "github.com/nathan-osman/george-the-dev-bot/reference"
	_ "github.com/nathan-osman/george-the-dev-bot/time"
)

func main() {
	var (
		addr     = flag.String("addr", ":8000", "HTTP address")
		email    = flag.String("email", os.Getenv("EMAIL"), "email address")
		password = flag.String("password", os.Getenv("PASSWORD"), "login password")
		room     = flag.Int("room", 1, "room to join")
	)
	flag.Parse()
	if *email == "" || *password == "" {
		log.Fatal("email and password cannot be blank")
	}
	s, err := server.New(*addr)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	c, err := sechat.New(*email, *password, *room)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connecting to chat...")
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	log.Print("Connected to chat")
	defer log.Print("Shutting down...")
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	for {
		select {
		case <-ch:
			return
		case e := <-c.Events:
			switch {
			// If the user is invited, have him join the room
			case e.EventType == sechat.EventInvitation:
				c.Join(e.RoomID)
			case e.IsMention:
				registry.Execute(c, e)
			}
		}
	}
}
