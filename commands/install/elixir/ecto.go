package elixir

import (
	"strings"

	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

const ectoSetupCommand = "mix ecto.setup"

func runEctoSetup() error {
	mixHelpStdout, _, err := runner.Run("mix help")
	if err != nil {
		return err
	}

	if strings.Contains(mixHelpStdout, "ecto.setup") {
		output.Found("Found ecto.setup Mix task. Running it to ensure database is setup:", 4)
		runner.StreamWithInfo(ectoSetupCommand, 6)
	}

	return nil
}
