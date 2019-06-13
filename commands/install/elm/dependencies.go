package elm

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func checkDependencies() {
	installCommand := "elm-package install --yes"
	output.Info("Checking all Elm dependencies are installed by running:", 2)
	runner.StreamWithInfo(installCommand, 2)
}
