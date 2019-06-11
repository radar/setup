package elixir

import (
	"errors"

	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func checkForErlang() error {
	erlangCheck := "which erl"
	output.Info("Elixir requires Erlang to be installed. Checking now...")
	output.Info("$ " + erlangCheck)
	return runner.OptionalAction(erlangCheck, erlangFound, erlangMissing)
}

func erlangFound() error {
	output.Found("Erlang is installed!")
	return nil
}

func erlangMissing() error {
	output.Fail("Erlang is missing!")
	output.Info("There are two ways that you can install Erlang.")
	output.Info("1) $ brew install erlang")
	output.Info("2) $ asdf plugin-add erlang && asdf install erlang <version>")
	output.Info("I wuold recommend the Homebrew way, as it will install the latest version.")

	return errors.New("Erlang is missing. Cannot install Elixir.")
}
