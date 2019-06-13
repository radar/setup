package elixir

import (
	"errors"

	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func checkForErlang() error {
	erlangCheck := "which erl"
	output.Info("Elixir requires Erlang to be installed. Checking now...", 4)
	output.Info("$ " + erlangCheck, 6)
	return runner.OptionalAction(erlangCheck, erlangFound, erlangMissing)
}

func erlangFound() error {
	output.Found("Erlang is installed!", 6)
	return nil
}

func erlangMissing() error {
	output.Fail("Erlang is missing!", 6)
	output.Info("There are two ways that you can install Erlang.", 8)
	output.Info("1) $ brew install erlang", 8)
	output.Info("2) $ asdf plugin-add erlang && asdf install erlang <version>", 8)
	output.Info("I wuold recommend the Homebrew way, as it will install the latest version.", 8)

	return errors.New("Erlang is missing. Cannot install Elixir.")
}
