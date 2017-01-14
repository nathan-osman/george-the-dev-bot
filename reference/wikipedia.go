package reference

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/nathan-osman/george-the-dev-bot/registry"
	"github.com/nathan-osman/go-sechat"
)

var (
	wikipediaRegexp = regexp.MustCompile(
		`(?i)^wh(?:at|o)\s+(?:was|is|were|are)\s+(.*[^?.])`,
	)
	wikipediaQueryURL   = "https://en.wikipedia.org/w/api.php?action=query&list=search&format=json&srsearch="
	wikipediaArticleURL = "https://en.wikipedia.org/wiki/"
)

// wikipediaResult represents the structure of data returned by the API.
type wikipediaResult struct {
	Query struct {
		Search []struct {
			Title string `json:"title"`
		} `json:"search"`
	} `json:"query"`
}

func init() {
	registry.Register(func(c *sechat.Conn, e *sechat.Event) {
		m := wikipediaRegexp.FindStringSubmatch(e.TextContent)
		if m != nil {
			// Reward anyone who asks "what is love"
			if strings.ToLower(m[1]) == "love" {
				c.Reply(e, "Baby don't hurt me. Don't hurt me no more!")
				return
			}
			go func() {
				r, err := http.Get(wikipediaQueryURL + url.QueryEscape(m[1]))
				if err != nil {
					c.Reply(e, "Hrm. Wikipedia is down. That's not good.")
					return
				}
				w := wikipediaResult{}
				if err := json.NewDecoder(r.Body).Decode(&r); err != nil ||
					len(w.Query.Search) == 0 || len(w.Query.Search[0].Title) == 0 {
					c.Reply(e, "Hrm. Wikipedia is returning malformed JSON.")
					return
				}
				c.Reply(e, wikipediaArticleURL+url.QueryEscape(w.Query.Search[0].Title))
			}()
		}
	})
}
