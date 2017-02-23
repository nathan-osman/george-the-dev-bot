package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/george-the-dev-bot/registry"
	"github.com/nathan-osman/george-the-dev-bot/server"
	"github.com/nathan-osman/go-sechat"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	_ "github.com/nathan-osman/george-the-dev-bot/apt"
	_ "github.com/nathan-osman/george-the-dev-bot/net"
	_ "github.com/nathan-osman/george-the-dev-bot/reference"
	_ "github.com/nathan-osman/george-the-dev-bot/time"
)

var log = logrus.New()

func main() {
	app := cli.NewApp()
	app.Name = "george"
	app.Usage = "run George the Dev bot"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "addr",
			Value:  ":8000",
			EnvVar: "ADDR",
			Usage:  "HTTP address",
		},
		cli.BoolFlag{
			Name:   "debug",
			EnvVar: "DEBUG",
			Usage:  "enable debugging",
		},
		cli.StringFlag{
			Name:   "email",
			EnvVar: "EMAIL",
			Usage:  "email address",
		},
		cli.StringFlag{
			Name:   "password",
			EnvVar: "PASSWORD",
			Usage:  "login password",
		},
		cli.IntFlag{
			Name:   "room",
			Value:  1,
			EnvVar: "ROOM",
			Usage:  "initial room to join",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		var (
			email    = c.String("email")
			password = c.String("password")
		)
		if email == "" || password == "" {
			log.Fatal("email and password cannot be blank")
		}
		srv, err := server.New(c.String("addr"))
		if err != nil {
			log.Fatal(err)
		}
		defer srv.Close()
		log.Print("initializing chat connection")
		conn, err := sechat.New(
			c.String("email"),
			c.String("password"),
			c.Int("room"),
		)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		defer log.Print("closing chat connection")
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		for {
			select {
			case <-ch:
				return nil
			case e := <-conn.Events:
				switch {
				// If the user is invited, have him join the room
				case e.EventType == sechat.EventInvitation:
					conn.Join(e.RoomID)
				case e.IsMention:
					registry.Execute(conn, e)
				}
			}
		}
	}
	app.Run(os.Args)
}
