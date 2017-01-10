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
	registry.Register(func(c *sechat.Conn, e *sechat.Event) {
		if uptimeRegexp.MatchString(e.Content) {
			c.Send(
				e.RoomID,
				fmt.Sprintf(
					":%d started %s",
					e.MessageID,
					timeago.FromTime(startTime),
				),
			)
		}
	})
}