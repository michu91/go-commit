package git

import (
	"os/exec"
	"strings"
)

func GetCurrentBranch() (string, error) {
	out := new(strings.Builder)
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Stdout = out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSuffix(out.String(), "\n"), nil
}
