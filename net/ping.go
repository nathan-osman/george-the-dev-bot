package net

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/nathan-osman/george-the-dev-bot/registry"
	"github.com/nathan-osman/george-the-dev-bot/util"
	"github.com/nathan-osman/go-sechat"
)

// The regexp is very conservative to avoid inadvertently triggering pings
var pingRegexp = regexp.MustCompile(
	`(?i)\bping(6)?\s+([\w-]+\.[\w.-]+)`,
)

func init() {
	registry.Register(func(c *sechat.Conn, e *sechat.Event) {
		m := pingRegexp.FindStringSubmatch(e.Content)
		if m != nil {
			go func() {
				cmd := exec.Command(
					fmt.Sprintf("ping%s", m[1]),
					"-c", "4",
					"-i", "0.2",
					"-w", "10",
					m[2],
				)
				o, _ := cmd.CombinedOutput()
				s := strings.TrimSpace(string(o))
				if len(s) == 0 {
					s = "<empty>"
				}
				c.Send(
					e.RoomID,
					util.Pre(s),
				)
			}()
		}
	})
}
