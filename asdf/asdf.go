package asdf

import (
	"strings"

	"github.com/radar/setup/runner"
)

type Tool struct {
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
