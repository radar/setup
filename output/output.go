package output

import (
	"github.com/fatih/color"
)

func Found(s string) {
	color.Blue("👀 " + s)
}

func Success(s string) {
	color.Green("✅ " + s)
}

func Fail(s string) {
	color.Red("❌ " + s)
}

func Info(s string) {
	color.Yellow("💡 " + s)
}

func Skip(s string) {
	color.HiBlack("👻 " + s)
}
