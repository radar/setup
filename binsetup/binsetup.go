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

	output.Found("Found a bin/setup script.")
	if (stat.Mode()&0111 == 0) {
		output.Fail("bin/setup is not executable!")
		output.Info("Make this file executable by running 'chmod +x bin/setup'.")
		return errors.New("bin/setup is not executable")
	}

	runner.Stream("bin/setup")
	return nil
}
