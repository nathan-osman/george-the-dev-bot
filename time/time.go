package time

import (
	"fmt"
	"regexp"
	"time"

	"github.com/nathan-osman/george-the-dev-bot/registry"
	"github.com/nathan-osman/go-sechat"
)

var timeRegexp = regexp.MustCompile(
	`(?i)current\s+(?:utc\s+)?time`,
)

func init() {
	registry.Register(func(c *sechat.Conn, e *sechat.Event) {
		if timeRegexp.MatchString(e.TextContent) {
			c.Reply(
				e,
				fmt.Sprintf(
					"the current time in UTC is %s",
					time.Now().UTC().Format("15:04:05"),
				),
			)
		}
	})
}
