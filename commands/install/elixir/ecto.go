package elixir

import (
	"strings"

	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

const ectoSetupCommand = "mix ecto.setup"

func runEctoSetup() error {
	mixHelpOutput, err := runner.Run("mix help")
	if err != nil {
		return err
	}

	if strings.Contains(mixHelpOutput, "ecto.setup") {
		output.Found("Found ecto.setup Mix task. Running it to ensure database is setup:")
		output.Info("$ " + ectoSetupCommand)
		runner.Stream(ectoSetupCommand)
	}

	return nil
}
