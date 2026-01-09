package logging

import (
	"fmt"
	"os"

	"github.com/mrumyantsev/mkpkg/internal/core"
)

func Info(text string) {
	fmt.Fprintln(os.Stdout, text)
}

func Infof(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, format, a...)
}

func Error(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", core.AppName, err.Error())
}

func ErrorBadquit(err error) {
	Error(err)
	badquit()
}

func ErrorHintGoodquit(err error) {
	Error(err)
	fmt.Fprintf(os.Stderr, core.Hint, core.AppName)
	goodquit()
}

func HelpGoodquit() {
	fmt.Fprintf(os.Stdout, core.Help, core.AppName)
	goodquit()
}

func VersionGoodquit() {
	fmt.Fprintf(os.Stdout, "%s version %s\n", core.AppName, core.Version)
	goodquit()
}

// goodquit exits the programm with status code 0.
func goodquit() {
	os.Exit(0)
}

// badquit exits the programm with status code 1.
func badquit() {
	os.Exit(1)
}
