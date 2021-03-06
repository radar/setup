package tool

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/radar/setup/asdf"
	"github.com/radar/setup/toolversions"
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/version"
)

type Tool struct {
	Name string
	PackageName string
	Executable string
	ExpectedVersion string
	VersionCommand string
	VersionRegexp string
}

func (tool *Tool) SetExpectedVersion() error {
	expectedVersion, err := toolversions.ForPackage(tool.PackageName)
	if err != nil {
		return err
	}

	output.FoundTitle(fmt.Sprintf("Found %s (%s) in .tool-versions", tool.Name, expectedVersion), 2)
	tool.ExpectedVersion = expectedVersion

	return nil
}

func (tool Tool) Install() error {
	tool.findExecutable()
	err := asdf.CheckInstallation(tool.PackageName, tool.ExpectedVersion)


	actualVersion, err := tool.actualVersion()
	if err != nil {
		return err
	}

	checker := version.Checker{
		tool.ExpectedVersion,
		actualVersion,
	}

	err = checker.Compare(tool.Name)
	if err != nil {
		return err
	}

	return nil
}

func (tool Tool) findExecutable() error {
	err := runner.LookPath(tool.VersionCommand)
	if err != nil {
		output.Fail(fmt.Sprintf("Could not find %s executable in PATH", tool.Executable), 4)
		return err
	}

	return nil
}

func (tool Tool) actualVersion() (string, error) {
	stdout, stderr, err := runner.Run(tool.VersionCommand)
	if err != nil {
		return stderr, err
	}

	re := regexp.MustCompile(tool.VersionRegexp)
	match := re.FindSubmatch([]byte(stdout))
	rubyVersion := strings.TrimSpace(string(match[1]))

	return rubyVersion, nil
}
