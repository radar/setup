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
		Remedy: installViaASDF,
	}

	err := tool.Install()
	if err != nil {
		return err
	}

	checkDependencies()

	return nil
}

func installViaASDF() string {
	return "To fix this issue, you can run \"asdf install\"."
}
