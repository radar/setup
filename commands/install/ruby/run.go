package ruby

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/radar/setup/common/version"
	"github.com/radar/setup/runner"
	"github.com/urfave/cli"
)

func Run(c *cli.Context) error {
	checker := version.Checker{
		Expected: expectedVersion(),
		Actual:   actualVersion(),
	}

	checker.Compare("Ruby", remedy)

	checkBundler()
	checkDependencies()
	return nil
}

func remedy() string {
	return "To fix this issue, you can run \"asdf install\"."
}

func expectedVersion() version.VersionCheckResult {
	dat, err := ioutil.ReadFile(".ruby-version")
	if err != nil {
		return version.VersionCheckResult{"", err}
	}

	rubyVersion := strings.TrimSpace(string(dat))

	return version.VersionCheckResult{rubyVersion, nil}
}

func actualVersion() version.VersionCheckResult {
	output, err := runner.Run("ruby -v")
	if err != nil {
		return version.VersionCheckResult{"", err}
	}

	re := regexp.MustCompile(`[\d\.]{3,}`)
	rubyVersion := strings.TrimSpace(string(re.Find([]byte(output))))

	return version.VersionCheckResult{rubyVersion, nil}
}
