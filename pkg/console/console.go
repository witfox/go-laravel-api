package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func Success(msg string) {
	colorOut(msg, "green")
}

func Error(msg string) {
	colorOut(msg, "red")
}

func Warning(msg string) {
	colorOut(msg, "yellow")
}

func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

//语法糖，判断err
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

func colorOut(msg string, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(msg, color))
}
