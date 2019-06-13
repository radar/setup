package binsetup

import (
	"errors"
	"os"

	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func RunIfExists() error {
	stat, err := os.Stat("bin/setup")

	if os.IsNotExist(err) {
		return nil
	}

	output.Found("Found a bin/setup script.", 2)

	if (stat.Size() == 0) {
		output.Info("But bin/setup is empty! I'll ignore it.", 4)
		return nil
	}

	if (stat.Mode()&0111 == 0) {
		output.Fail("bin/setup is not executable!", 4)
		output.Info("Make this file executable by running 'chmod +x bin/setup'.", 6)
		return errors.New("bin/setup is not executable")
	}


	runner.Stream("bin/setup", 4)
	return nil
}
