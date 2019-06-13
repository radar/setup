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
	output.FoundTitle("This looks like a go project!", 2)
	output.Info("Running: go build", 4)
	runner.Stream("go build", 6)
	return nil
}
