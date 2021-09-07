package exec

import (
	"os"
	"os/exec"
)

// getShell returns the current users shell, if no shell is found fallback to /bin/sh.
func getShell() string {
	shell, ok := os.LookupEnv("SHELL")
	if !ok {
		return "/bin/sh"
	}

	return shell
}

// RunCommand executes a command prints the output to stdout.
func RunCommand(command string) (string, error) {
	shell := getShell()

	output, err := exec.Command(shell, "-c", command).Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}