package version

import (
	"errors"
	"github.com/radar/setup/output"
	"fmt"
)

type Checker struct {
	Expected, Actual string
}

func (c Checker) Compare(name string) error {
	if c.equal() {
		output.Success(fmt.Sprintf("Correct %s version installed (%s)", name, c.Expected), 6)
	} else {
		output.Fail(fmt.Sprintf("Incorrect %s version installed: %s, was expecting %s", name, c.Actual, c.Expected), 6)
		return errors.New("Version comparison failed")
	}

	return nil
}
func (c Checker) equal() bool {
	return c.Expected == c.Actual
}
