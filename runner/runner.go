package runner

import (
	"fmt"
	"strings"
	"os/exec"
	"bufio"
	"bytes"

	"github.com/radar/setup/output"
)

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

func LookPath(command string) error {
	parts := strings.Split(command, " ")
	_, err := exec.LookPath(parts[0])
	return err
}

func buildCommand(command string) *exec.Cmd {
	parts := strings.Split(command, " ")
	return exec.Command(parts[0], parts[1:]...)
}
