package homebrew

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/utils"
)

func Bundle() error {
	if utils.FileExists("Brewfile") {
		output.Found("Found a Brewfile!", 0)
		output.Info("Checking Brewfile dependencies are installed...", 2)
		installBundle()
	}

	return nil
}

func installBundle() {
	runner.OptionalAction("brew bundle check", bundleInstalled, bundleNotInstalled)
}

func bundleInstalled() error {
	output.Success("Homebrew bundle installed.", 4)
	return nil
}

func bundleNotInstalled() error {
	output.Info("Homebrew bundle not installed. You might want to check the output of 'brew bundle check'", 4)
	output.Info("Alternatively, you could run 'brew bundle install' to install the dependencies listed in the Brewfile", 4)

	return nil
}
