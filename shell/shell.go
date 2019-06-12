package shell

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

type Shell string
const (
	ZSH Shell = "zsh"
	Bash Shell = "bash"
	Unknown Shell = "unknown"
)

func Run(command string) (stdout bytes.Buffer, stderr bytes.Buffer, err error) {
	var (
		cmdName string
		cmdArgs []string
	)

	switch detectShell() {
	case ZSH:
		cmdName = "zsh"
		cmdArgs = []string{"-c", ". ~/.zshrc; " + command}
	default:
		err = errors.New("Unable to detect shell! ")
		return
	}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()

	return
}

type action func() error

func OptionalAction(command string, success action, remedy action) (err error) {
	_, _, err = Run(command)
	if err != nil {
		return remedy()
	}

	return success()
}

func detectShell() Shell {
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "zsh") {
		return ZSH
	} else if strings.Contains(shell, "bash") {
		return Bash
	}

	return Unknown
}
