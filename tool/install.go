package tool

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/radar/setup/asdf"
	"github.com/radar/setup/nvm"
	"github.com/radar/setup/toolversions"
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/version"
)

type Source int

const (
	ASDF   Source = 0
	NVM    Source = 1
)

func (source Source) String() string {
	switch source {
	case ASDF:
		return "asdf"
	case NVM:
		return "NVM"
	}

	return ""
}

type Tool struct {
	Name string
	PackageName string
	Executable string
	ExpectedVersion string
	VersionCommand string
	VersionRegexp string
	Sources []Source
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

func (tool Tool) Install() (err error) {
	err = tool.findExecutable()
	installationSuccess := false

	if err != nil {
		for _, source := range tool.Sources {
			err = tool.checkInstallation(source)
			if err == nil {
				installationSuccess = true
				break
			}
		}

		if err != nil {
			return err
		}
	} else {
		installationSuccess = true
	}

	if (!installationSuccess) {
		panic("Failed to install!")
	}

	actualVersion, err := tool.actualVersion()
	if err != nil {
		fmt.Println(actualVersion)
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

func (tool Tool) checkInstallation(source Source) (err error) {
	switch source {
	case ASDF:
		err = asdf.CheckInstallation(tool.PackageName, tool.ExpectedVersion)
	case NVM:
		if nvm.Present() {
			output.Found("NVM is installed. I will use that to install Node.")
			err = nvm.CheckInstallation(tool.ExpectedVersion)
		}
	default:
		panic("WHAT DO I DO WITH "+ source.String())
	}

	return err
}

func (tool Tool) actualVersion() (string, error) {
	stdout, stderr, err := shell.Run(tool.VersionCommand)
	if err != nil {
		return stderr, err
	}

	re := regexp.MustCompile(tool.VersionRegexp)
	match := re.FindSubmatch([]byte(stdout))
	version := strings.TrimSpace(string(match[1]))

	return version, nil
}
