package runner

import (
	"fmt"
	"strings"
	"os/exec"
	"bufio"
	"bytes"
)

func Run(command string) (output string, err error) {
	cmd := exec.Command("elixir", "-v")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()

	return out.String(), err
}

func Stream(command string) {
	parts := strings.Split(command, " ")

	cmd := exec.Command(parts[0], parts[1:]...)
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
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(out.String())
		panic(err)
	}

	if strings.Contains(out.String(), text) {
		remedy()
	} else {
		success()
	}
}
