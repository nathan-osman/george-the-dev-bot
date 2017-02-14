package time

import (
	"fmt"
	"regexp"
	"time"

	"github.com/justincampbell/timeago"
	"github.com/nathan-osman/george-the-dev-bot/registry"
	"github.com/nathan-osman/go-sechat"
)

var (
	uptimeRegexp = regexp.MustCompile(`(?i)\buptime\b`)
	startTime    = time.Now()
)

func init() {
	registry.Register(func(c *sechat.Conn, e *sechat.Event) bool {
		if uptimeRegexp.MatchString(e.TextContent) {
			c.Reply(
				e,
				fmt.Sprintf(
					"started %s",
					timeago.FromTime(startTime),
				),
			)
			return true
		}
		return false
	}, registry.RegularCommand)
}
