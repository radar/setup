package golang

import (
	"os"

	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func Used() bool {
	return checkForGoMod() || checkForMain()
}

func checkForGoMod() bool {
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return false
	}

	return true
}

func checkForMain() bool {
	if _, err := os.Stat("main.go"); os.IsNotExist(err) {
		return false
	}

	return true
}

func Run() error {
	output.Found("This looks like a go project!")
	output.Info("Running: go build")
	runner.Stream("go build")
	return nil
}
