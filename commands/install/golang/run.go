package golang

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/utils"
)

func Used() bool {
	return utils.FileExists("go.mod") || utils.FileExists("main.go")
}

func Run() error {
	output.Found("This looks like a go project!")
	output.Info("Running: go build")
	runner.Stream("go build")
	return nil
}
