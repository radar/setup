package tool

import (
	"errors"
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
	Remedy (func() string)
}

func (tool Tool) Install() error {
	expectedVersion, err := toolversions.ForPackage(tool.PackageName)
	if err != nil {
		return err
	}

	output.Found(fmt.Sprintf("Found %s (%s) in .tool-versions", tool.Name, expectedVersion))
	tool.ExpectedVersion = expectedVersion

	err = tool.findExecutable()
	if err != nil {
		return err
	}

	err = tool.ensureInstalled(false)
	if err != nil {
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

func (tool Tool) findExecutable() error {
	err := runner.LookPath(tool.VersionCommand)
	if err != nil {
		output.Fail(fmt.Sprintf("Could not find % executable in PATH", tool.Executable))
		output.Info(tool.Remedy())
		return err
	}

	return nil
}

func (tool Tool) ensureInstalled(attempted bool) error {
	asdfTool := asdf.ListVersions(tool.PackageName)
	if asdfTool.CheckInstalled(tool.ExpectedVersion) {
		return nil
	}

	errorMsg := fmt.Sprintf("You do not have %s (%s) installed.", tool.Name, tool.ExpectedVersion)
	output.Fail(errorMsg)
	if (attempted) {
		output.Fail("Prior installation attempt failed. Please try it yourself with 'asdf install'")
		return errors.New(fmt.Sprintf("Could not install %s (%s)", tool.Name, tool.ExpectedVersion))
	}
	asdfTool.Install(tool.ExpectedVersion)

	return 	tool.ensureInstalled(true)
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
