package node

import (
	"github.com/radar/setup/tool"
)

func Run() error {
	tool := tool.Tool{
		Name: "Node",
		PackageName: "nodejs",
		Executable: "node",
		VersionCommand: "node -v",
		VersionRegexp: `([\d\.]{3,})`,
	}

	err := tool.Install()
	if err != nil {
		return err
	}

	checkDependencies()

	return nil
}
