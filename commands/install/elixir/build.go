package elixir

import (
	"github.com/radar/setup/runner"
)


func ensureHexPresent() {
	runner.StreamWithInfo("mix local.hex --force")
}

func ensureRebarPresent() {
	runner.StreamWithInfo("mix local.rebar --force")
}
