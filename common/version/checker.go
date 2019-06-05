package version

import (
	"github.com/radar/setup/output"
	"fmt"
)

type VersionCheckResult struct {
	Version string
	Err     error
}

type Checker struct {
	Expected, Actual VersionCheckResult
}

type remedy func() (string)

func (c Checker) Compare(name string, remedy remedy) {
	if c.equal() {
	  output.Success(fmt.Sprintf("Correct %s version installed (%s)", name, c.Expected.Version))
	} else {
		output.Fail(fmt.Sprintf("Incorrect %s version installed: %s, was expecting %s", name, c.Actual.Version, c.Expected.Version))
		output.Info(remedy())
	}
}
func (c Checker) equal() bool {
	return c.Expected.Version == c.Actual.Version
}
