package ruby

import (
	"github.com/radar/setup/tool"
)


func checkRuby() (err error) {
	tool := tool.Tool{
		Name: "Ruby",
		PackageName: "ruby",
		Executable: "ruby",
		VersionCommand: "ruby -v",
		VersionRegexp: `([\d\.]{3,})`,
	}

	err = tool.SetExpectedVersion()
	if err != nil {
		return err
	}

	return tool.Install()
}
