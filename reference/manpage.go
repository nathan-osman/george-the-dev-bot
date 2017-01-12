package reference

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/nathan-osman/george-the-dev-bot/registry"
	"github.com/nathan-osman/go-sechat"
)

var manpageRegexp = regexp.MustCompile(
	`(?i)(?:^man\s+([\w-]+)|manpage\s+(?:for\s+(?:the\s+)?)?([\w-]+)|([\w-]+)\s+manpage)`,
)

func init() {
	registry.Register(func(c *sechat.Conn, e *sechat.Event) {
		m := manpageRegexp.FindStringSubmatch(e.Content)
		if m != nil {
			go func() {
				// Grab the subexpressions that matched
				var (
					item = m[1] + m[2] + m[3]
					url  = "http://manpages.ubuntu.com/" + p
				)
				r, err := http.Get(url)
				if err != nil || r.StatusCode >= 400 {
					c.Reply(e, "unable to find manpage")
				}
				c.Reply(e, fmt.Sprintf(
					"[manpage for %s](%s)",
					item,
					url,
				))
			}()
		}
	})
}
