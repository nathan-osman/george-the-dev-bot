package net

import (
	"regexp"

	"github.com/nathan-osman/george-the-dev-bot/registry"
	"github.com/nathan-osman/george-the-dev-bot/util"
	"github.com/nathan-osman/go-sechat"
)

var digRegexp = regexp.MustCompile(
	`(?i)\bdig\s+(?:(a|a{4}|mx)\s+)?([\w-]+\.[\w.-]+)`,
)

func init() {
	registry.Register(func(c *sechat.Conn, e *sechat.Event) {
		m := digRegexp.FindStringSubmatch(e.TextContent)
		if m != nil {
			go func() {
				if len(m[1]) == 0 {
					m[1] = "a"
				}
				c.Send(
					e.RoomID,
					util.Exec(
						"dig",
						m[1],
						m[2],
					),
				)
			}()
		}
	})
}
