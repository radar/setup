package ruby

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"

	"github.com/radar/setup/common/version"
	"github.com/urfave/cli"
)

func Run(c *cli.Context) error {
	checker := version.Checker{
		Expected: expectedVersion(),
		Actual:   actualVersion(),
	}

	checker.Compare("Ruby", remedy)
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
	cmd := exec.Command("ruby", "-v")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return version.VersionCheckResult{"", err}
	}

	re := regexp.MustCompile(`[\d\.]{3,}`)
	rubyVersion := strings.TrimSpace(string(re.Find([]byte(out.String()))))

	return version.VersionCheckResult{rubyVersion, nil}
}
