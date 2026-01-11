package config

import (
	"errors"
	"flag"
	"io"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/mrumyantsev/mkpkg/internal/logging"
	"github.com/stoewer/go-strcase"
)

type Config struct {
	Path              string
	Filename          string
	Eol               string
	PackageName       string
	ObjectName        string
	ObjectType        string
	ObjectConstructor string
	ObjectReceiver    string
	ObjectMethods     []string
}

func New() *Config {
	return &Config{
		Eol:               "\\n",
		ObjectType:        "struct",
		ObjectConstructor: "New()",
	}
}

func (c *Config) ParseCliArgs(args []string) {
	isDebug := false

	for i := range args {
		if args[i] == "--dbg" {
			isDebug = true
		}
		if args[i] == "-v" || args[i] == "--version" {
			logging.VersionGoodquit()
		}
		if args[i] == "-iface" || args[i] == "--iface" {
			c.ObjectType = "interface"
		}
	}

	args = removeFromArgs(args, "-iface", "--iface", "--dbg")

	flag.CommandLine.Init("", flag.ContinueOnError)
	flag.CommandLine.SetOutput(io.Discard)

	flagC := flag.String("c", "", "")
	flagCtor := flag.String("ctor", "", "")
	flagEol := flag.String("eol", "", "")
	flagF := flag.String("f", "", "")
	flagFilename := flag.String("filename", "", "")
	flagM := flag.String("m", "", "")
	flagMethods := flag.String("methods", "", "")
	flagN := flag.String("n", "", "")
	flagName := flag.String("name", "", "")
	flagP := flag.String("p", "", "")
	flagPackage := flag.String("package", "", "")

	err := flag.CommandLine.Parse(args[1:])
	if errors.Is(err, flag.ErrHelp) {
		logging.HelpGoodquit()
	}
	if err != nil {
		logging.ErrorHintGoodquit(err)
	}

	c.Filename = pickOneVal(c.Filename, flagF, flagFilename)
	c.Eol = pickOneVal(c.Eol, flagEol)
	c.Eol = strings.ReplaceAll(c.Eol, "\\n", "\n")
	c.Eol = strings.ReplaceAll(c.Eol, "\\r", "\r")
	c.PackageName = pickOneVal(c.PackageName, flagP, flagPackage)
	c.ObjectName = pickOneVal(c.ObjectName, flagN, flagName)
	c.ObjectConstructor = pickOneVal(c.ObjectConstructor, flagC, flagCtor)

	objectMethodsStr := pickOneVal("", flagM, flagMethods)

	if objectMethodsStr != "" {
		c.ObjectMethods = strings.Split(objectMethodsStr, ";")
	}

	var public, private []string
	var method string

	for i := range c.ObjectMethods {
		method = strings.TrimSpace(c.ObjectMethods[i])

		if len(method) >= 1 && method[0] >= 'A' && method[0] <= 'Z' {
			public = append(public, method)
		} else {
			private = append(private, method)
		}
	}

	c.ObjectMethods = append(public, private...)

	nonflags := flag.Args()

	switch len(nonflags) {
	case 0:
		logging.ErrorHintGoodquit(errors.New("package path not specified"))
	case 1:
		c.Path = nonflags[0]
	default:
		logging.ErrorHintGoodquit(errors.New("multiple package paths are not allowed"))
	}

	basename := filepath.Base(c.Path)

	if c.Filename == "" {
		c.Filename = basename
	}

	if len(c.Filename) >= 4 && c.Filename[len(c.Filename)-3:] != ".go" ||
		len(c.Filename) <= 3 {
		c.Filename += ".go"
	}

	if c.PackageName == "" {
		c.PackageName = strings.ReplaceAll(basename, "-", "")
		c.PackageName = strings.ToLower(c.PackageName)
	}

	if c.ObjectName == "" {
		c.ObjectName = strcase.UpperCamelCase(c.Filename[:len(c.Filename)-3])
	}

	if len(c.ObjectName) >= 1 {
		c.ObjectReceiver = c.ObjectName[:1]
	}

	if len(c.ObjectReceiver) >= 1 && unicode.IsUpper(rune(c.ObjectReceiver[0])) {
		c.ObjectReceiver = strings.ToLower(c.ObjectReceiver)
	}

	if isDebug {
		c.showDebug()
	}
}

func pickOneVal(def string, vals ...*string) string {
	for i := range vals {
		if vals[i] != nil && *vals[i] != "" {
			return *vals[i]
		}
	}

	return def
}

func removeFromArgs(args []string, remove ...string) []string {
	removed := 0

	for i := range args {
		if isTargetArg(args[i], remove...) {
			removed++

			continue
		}

		args[i-removed] = args[i]
	}

	return args[:len(args)-removed]
}

func isTargetArg(arg string, targets ...string) bool {
	for i := range targets {
		if targets[i] == arg {
			return true
		}
	}

	return false
}
