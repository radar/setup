package elixir

import (
	"regexp"
	"strings"

	"github.com/radar/setup/runner"
	"github.com/radar/setup/common/toolversions"
	"github.com/radar/setup/common/version"
)

func Run() error {
	checker := version.Checker{
		Expected: expectedVersion(),
		Actual:   actualVersion(),
	}

	checker.Compare("Elixir", installViaASDF)
	checkDependencies()

	return nil
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
