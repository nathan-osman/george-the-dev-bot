package apt

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nathan-osman/george-the-dev-bot/registry"
	"github.com/nathan-osman/george-the-dev-bot/util"
	"github.com/nathan-osman/go-sechat"
)

var aptRegexp = regexp.MustCompile(
	`(?i)\bwh(?:at|ich)\s+package\s+provides\s+(?:the\s+file\s+)?([\w/.-]+)`,
)

func init() {
	registry.Register(func(c *sechat.Conn, e *sechat.Event) {
		m := aptRegexp.FindStringSubmatch(e.TextContent)
		if m != nil {
			go func() {
				var (
					o = util.Exec(
						"apt-file",
						"search",
						m[1],
					)
					l = strings.Split(o, "\n")
					p = strings.Split(o, ":")
				)
				if len(l) > 1 {
					c.Reply(e, fmt.Sprintf("%d packages provide this file", len(l)))
					return
				}
				if len(p) == 1 {
					c.Reply(e, fmt.Sprintf("no packages provide %s", m[1]))
					return
				}
				c.Reply(e, fmt.Sprintf("the %s package provides %s", p[0], m[1]))
			}()
		}
	})
}
