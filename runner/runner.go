package runner

import (
	"strings"
	"os/exec"
	"bufio"
	"bytes"

	"github.com/radar/setup/output"
)

func StreamWithInfo(command string, padding output.Padding) {
	output.Info("$ " + command, padding)
	Stream(command, padding)
}

func Run(command string) (string, string, error) {
	cmd := buildCommand(command)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return stdout.String(), stderr.String(), err
		}
	}

	return stdout.String(), stderr.String(), nil
}

func Stream(command string, padding output.Padding) {
	cmd := buildCommand(command)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		output.Info(m, padding)
	}
	cmd.Wait()
}

type action func() error

func CheckForMessage(command string, text string, success action, remedy action) {
	stdout, _, err := attemptCommand(command)
	if err != nil {
		output.Fail("Command failed: " + command, 0)
		output.Info("Attempting a remedy...", 0)
		remedy()
	}

	if strings.Contains(stdout.String(), text) {
		remedy()
	} else {
		success()
	}
}

func OptionalAction(command string, success action, fail action) error {
	_, _, err := attemptCommand(command)
	if err != nil {
		return fail()
	} else {
		return success()
	}
}

func LookPath(command string) error {
	parts := strings.Split(command, " ")
	_, err := exec.LookPath(parts[0])
	return err
}

func buildCommand(command string) *exec.Cmd {
	parts := strings.Split(command, " ")
	return exec.Command(parts[0], parts[1:]...)
}

func attemptCommand(command string) (stdout bytes.Buffer, stderr bytes.Buffer, err error) {
	parts := strings.Split(command, " ")

	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	return stdout, stderr, err
}
