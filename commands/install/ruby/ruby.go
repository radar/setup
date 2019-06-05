package ruby

import (
	"github.com/radar/setup/tool"
)


func checkRuby() error {
	tool := tool.Tool{
		Name: "Ruby",
		PackageName: "ruby",
		Executable: "ruby",
		VersionCommand: "ruby -v",
		VersionRegexp: `([\d\.]{3,})`,
		Remedy: installViaASDF,
	}

	return tool.Install()
}

func installViaASDF() string {
	return "To fix this issue, you can run \"asdf install\"."
}
