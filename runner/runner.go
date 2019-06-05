package runner

import (
	"fmt"
	"strings"
	"os/exec"
	"bufio"
	"bytes"

	"github.com/radar/setup/output"
)

func Run(command string) (output string, err error) {
	cmd := buildCommand(command)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println(exitError.ExitCode())
		}
	}

	return out.String(), err
}

func Stream(command string) {
	cmd := buildCommand(command)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
	}
	cmd.Wait()
}

type action func()

func CheckForMessage(command string, text string, success action, remedy action) {
	parts := strings.Split(command, " ")

	cmd := exec.Command(parts[0], parts[1:]...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		output.Fail("Command failed: " + command)
		output.Info("Attempting a remedy...")
		remedy()
	}

	if strings.Contains(stdout.String(), text) {
		remedy()
	} else {
		success()
	}
}

func buildCommand(command string) *exec.Cmd {
	parts := strings.Split(command, " ")
	return exec.Command(parts[0], parts[1:]...)
}
