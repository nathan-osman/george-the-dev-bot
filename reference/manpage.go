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
		m := manpageRegexp.FindStringSubmatch(e.TextContent)
		if m != nil {
			go func() {
				// Grab the subexpressions that matched
				var (
					item = m[1] + m[2] + m[3]
					url  = "http://manpages.ubuntu.com/cgi-bin/search.py?titles=404&q=" + item
				)
				r, err := http.Get(url)
				if err != nil || r.StatusCode >= 400 {
					c.Reply(e, fmt.Sprintf(
						"unable to find manpage for %s.",
						item,
					))
					return
				}
				c.Reply(e, fmt.Sprintf(
					"[manpage for %s](%s)",
					item,
					r.Request.URL.String(),
				))
			}()
		}
	})
}
