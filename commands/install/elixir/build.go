package elixir

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)


func ensureHexPresent() {
	output.Info("Ensuring Hex is present", 4)
	runner.StreamWithInfo("mix local.hex --force", 6)
}

func ensureRebarPresent() {
	output.Info("Ensuring Rebar is present", 4)
	runner.StreamWithInfo("mix local.rebar --force", 6)
}
