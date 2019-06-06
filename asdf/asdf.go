package asdf

import (
	"fmt"
	"strings"

	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

type Tool struct {
	Name string
	Versions []string
}

func ListVersions(tool string) Tool {
	listOutput, _, _ := runner.Run("asdf list " + tool)
	rawVersions := strings.Split(strings.TrimSpace(listOutput), "\n")

	versions := make([]string, len(rawVersions))
	for i, v := range rawVersions {
			versions[i] = strings.TrimSpace(v)
	}

	return Tool{
		Name: tool,
		Versions: versions,
	}
}

func (t Tool) CheckInstalled(expectedVersion string) bool {
    for _, actualVersion := range t.Versions {
        if expectedVersion == actualVersion {
            return true
        }
    }
    return false
}

func (t Tool) Install(version string ) {
	installCommand := fmt.Sprintf("asdf install %s %s", t.Name, version)
	output.Info("Attempting installation:")
	output.Info("$ " + installCommand)
	runner.Stream(installCommand)
}
