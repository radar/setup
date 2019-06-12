package nvm

import (
	"fmt"

	"github.com/radar/setup/output"
	"github.com/radar/setup/shell"
)

type NVM struct {
	ExpectedVersion string
}

func Present() bool {
	_, _, err := shell.Run("which nvm")
	if err != nil {
		return false
	}

	return true
}

func CheckInstallation(expectedVersion string) (err error) {
	nvm := NVM{
		ExpectedVersion: expectedVersion,
	}
	err = nvm.checkVersionInstalled()
	if err != nil {
		return err
	}

	return nil
}

func (nvm NVM) checkVersionInstalled() (err error) {
	return shell.OptionalAction(
		"nvm version " + nvm.ExpectedVersion,
		nvm.correctVersionInstalled,
		nvm.installExpectedVersion,
	)
}

func (nvm NVM) correctVersionInstalled() (err error) {
	return nil
}

func (nvm NVM) installExpectedVersion() (err error) {
	output.Info(fmt.Sprintf("Installing Node (%s) with NVM...", nvm.ExpectedVersion))
	installCmd := "nvm install " + nvm.ExpectedVersion
	output.Info("$ " + installCmd)
	stdout, stderr, err := shell.Run(installCmd)
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())

	return err

}
