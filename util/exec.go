package util

import (
	"os/exec"
	"strings"
)

// Exec runs the specified command and returns its output.
func Exec(name string, args ...string) string {
	c := exec.Command(name, args...)
	o, _ := c.CombinedOutput()
	s := strings.TrimSpace(string(o))
	if len(s) == 0 {
		return Pre("<no output>")
	}
	return Pre(s)
}
