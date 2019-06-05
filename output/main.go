package output

import (
	"github.com/fatih/color"
)

func Success(s string) {
	color.Green("✅ " + s)
}

func Fail(s string) {
	color.Red("❌ " + s)
}

func Info(s string) {
	color.Yellow("💡 " + s)
}
