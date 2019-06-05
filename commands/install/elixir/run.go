package elixir

import (
	"regexp"
	"strings"

	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/common/toolversions"
	"github.com/radar/setup/common/version"
	"github.com/urfave/cli"
)

func Run(c *cli.Context) error {
	checker := version.Checker{
		Expected: expectedVersion(),
		Actual:   actualVersion(),
	}

	checker.Compare("Elixir", installViaASDF)

	output.Info("Checking all dependencies are installed by running 'mix deps'...")
	runner.CheckForMessage(
		"mix deps",
		"the dependency is not available",
		dependenciesInstalled,
		installDependencies,
	)

	return nil
}

func dependenciesInstalled() {
	output.Success("Elixir dependencies are installed.")
}

func installDependencies() {
	output.Fail("Hex packages are missing.")
	output.Info("Attempting installation with:")
	output.Info("$ mix hex.local --if-missing")
	output.Info("$ mix deps.get")

	runner.Stream("mix hex.local --if-missing")
	runner.Stream("mix deps.get")
}

func installViaASDF() string {
	return "To fix this issue, you can run \"asdf install\"."
}

func expectedVersion() version.VersionCheckResult {
	result, err := toolversions.ForPackage("elixir")
	return version.VersionCheckResult{result, err}
}

func actualVersion() version.VersionCheckResult {
	output, err := runner.Run("elixir -v")
	if err != nil {
		return version.VersionCheckResult{"", err}
	}

	re := regexp.MustCompile(`Elixir ([\d+\.]{3,})`)
  match := re.FindSubmatch([]byte(output))
	actualVersion := strings.TrimSpace(string(match[1]))

	return version.VersionCheckResult{actualVersion, nil}
}
