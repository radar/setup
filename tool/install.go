package tool

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/radar/setup/toolversions"
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/version"
)

type Tool struct {
	Name string
	PackageName string
	Executable string
	VersionCommand string
	VersionRegexp string
	Remedy (func() string)
}

func (tool Tool) Install() error {
	expectedVersion, err := toolversions.ForPackage(tool.PackageName)
	if err != nil {
		return err
	}

	output.Found(fmt.Sprintf("Found %s (%s) in .tool-versions", tool.Name, expectedVersion))

	err = runner.LookPath(tool.VersionCommand)
	if err != nil {
		output.Fail(fmt.Sprintf("Could not find % executable in PATH", tool.Executable))
		output.Info(tool.Remedy())
		return err
	}

	actualVersion, err := tool.actualVersion()
	if err != nil {
		return err
	}

	checker := version.Checker{
		expectedVersion,
		actualVersion,
	}
	err = checker.Compare(tool.Name, tool.Remedy)
	if err != nil {
		return err
	}

	return nil
}

func (tool Tool) actualVersion() (string, error) {
	output, err := runner.Run(tool.VersionCommand)
	if err != nil {
		return output, err
	}

	re := regexp.MustCompile(tool.VersionRegexp)
	match := re.FindSubmatch([]byte(output))
	rubyVersion := strings.TrimSpace(string(match[1]))

	return rubyVersion, nil
}
