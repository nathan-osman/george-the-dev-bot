package util

import (
	"strings"
)

// indent is inserted before each line.
var indent = strings.Repeat(" ", 4)

// Pre adds four spaces to the beginning of each line to ensure that it is
// displayed in a fixed width font.
func Pre(s string) string {
	return indent + strings.Replace(s, "\n", "\n"+indent, -1)
}
