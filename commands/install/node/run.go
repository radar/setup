package node

import (
	"github.com/radar/setup/tool"
)

func Run() (err error) {
	tool := tool.Tool{
		Name: "Node",
		PackageName: "nodejs",
		Executable: "node",
		VersionCommand: "node -v",
		VersionRegexp: `([\d\.]{3,})`,
	}


	err = tool.SetExpectedVersion()
	if err != nil {
		return err
	}

	err = tool.Install()
	if err != nil {
		return err
	}

	checkDependencies()

	return nil
}
