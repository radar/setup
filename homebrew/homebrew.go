package homebrew

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/utils"
)

func Bundle() error {
	if utils.FileExists("Brewfile") {
		output.Found("Found a Brewfile!")
		output.Info("Checking Brewfile dependencies are installed...")
		installBundle()
	}

	return nil
}

func installBundle() {
	runner.OptionalAction("brew bundle check", bundleInstalled, bundleNotInstalled)
}

func bundleInstalled() error {
	output.Success("Homebrew bundle installed.")
	return nil
}

func bundleNotInstalled() error {
	output.Info("Homebrew bundle not installed. Installing...")
	runner.StreamWithInfo("brew bundle install -v")
	installBundle()

	return nil
}
