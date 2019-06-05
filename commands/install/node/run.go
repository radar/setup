package node

import (
	"bytes"

	"os/exec"
	"regexp"
	"strings"

	"github.com/radar/setup/common/toolversions"
	"github.com/radar/setup/common/version"
)

func Run() error {
	checker := version.Checker{
		Expected: expectedVersion(),
		Actual:   actualVersion(),
	}

	checker.Compare("Node", remedy)
	checkDependencies()
	return nil
}

func remedy() string {
	return "To fix this issue, you can run \"asdf install\"."
}

func expectedVersion() version.VersionCheckResult {
	result, err := toolversions.ForPackage("nodejs")
	return version.VersionCheckResult{result, err}
}

func actualVersion() version.VersionCheckResult {
	cmd := exec.Command("node", "-v")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return version.VersionCheckResult{"", err}
	}

	re := regexp.MustCompile(`[\d+\.]{3,}`)
	actualVersion := strings.TrimSpace(string(re.Find([]byte(out.String()))))

	return version.VersionCheckResult{actualVersion, nil}
}
