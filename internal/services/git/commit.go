package git

import (
	"io"
	"os"
	"os/exec"
)

func CommitCommand(message string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}
	bytes, err := io.ReadAll(stdOut)
	if err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return string(bytes), nil
}
