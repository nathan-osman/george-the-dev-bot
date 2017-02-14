package net

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/nathan-osman/george-the-dev-bot/registry"
	"github.com/nathan-osman/go-sechat"
)

var isDownRegexp = regexp.MustCompile(
	`(?i)\bis\s+([\w.-]+)\s+down\b`,
)

func init() {
	registry.Register(func(c *sechat.Conn, e *sechat.Event) bool {
		m := isDownRegexp.FindStringSubmatch(e.TextContent)
		if m != nil {
			go func() {
				r, err := http.Get(
					fmt.Sprintf("http://%s", m[1]),
				)
				if err != nil || r.StatusCode >= 400 {
					c.Reply(e, fmt.Sprintf(
						"%s seems to be down",
						m[1],
					))
					return
				}
				c.Reply(e, fmt.Sprintf(
					"%s works for me",
					m[1],
				))
			}()
			return true
		}
		return false
	}, registry.RegularCommand)
}
