package elm

import (
	"github.com/radar/setup/tool"
)

func Run() error {
	tool := tool.Tool{
		Name: "Elm",
		PackageName: "elm",
		Executable: "elm",
		VersionCommand: "elm -v",
		VersionRegexp: `([\d\.]{3,})`,
	}

	err := tool.Install()
	if err != nil {
		return err
	}

	checkDependencies()

	return nil
}
