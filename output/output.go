package output

import (
	"fmt"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/fatih/color"
)

type Padding int

// DefaultInitialPadding is the default padding in the log library.
const DefaultInitialPadding Padding = 0

func Found(s string, padding Padding) {
	blue := color.New(color.FgBlue)
	info(blue.Sprint("👀 " + s), padding)
}

func FoundTitle(s string, padding Padding) {
	boldBlue := color.New(color.FgBlue, color.Bold)
	info(boldBlue.Sprint("👀 " + s), padding)
}

func Success(s string, padding Padding) {
	green := color.New(color.FgGreen)
	info(green.Sprint("✅ " + s), padding)
}

func Fail(s string, padding Padding) {
	red := color.New(color.FgRed)
	info(red.Sprint("❌ " + s), padding)
}

func Info(title string, padding Padding) {
	defer func() {
		cli.Default.Padding = int(DefaultInitialPadding)
	}()
	cli.Default.Padding = int(padding)
	log.Infof(color.New(color.Bold).Sprint(title))
	cli.Default.Padding = int(padding + DefaultInitialPadding)
}

func info(s string, padding Padding) {
	defer func() {
		cli.Default.Padding = int(DefaultInitialPadding)
	}()
	cli.Default.Padding = int(padding)
	log.Infof(s)
	cli.Default.Padding = int(padding + DefaultInitialPadding)
}

func Skip(s string, padding Padding) {
	black := color.New(color.FgHiBlack)
	info(black.Sprint("👻 " + s), padding)
}

func Separator() {
	fmt.Println(" ")
}
