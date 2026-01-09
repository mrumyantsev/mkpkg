package mkpkg

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mrumyantsev/mkpkg/internal/config"
	"github.com/mrumyantsev/mkpkg/internal/logging"
)

type App struct {
	config *config.Config
}

func New() *App {
	cfg := config.New()

	cfg.ParseCliArgs(os.Args)

	return &App{
		config: cfg,
	}
}

func (a *App) Run() {
	makeAllDirs(a.config.Path)

	fpath := filepath.Join(a.config.Path, a.config.Filename)

	if !isFileExists(fpath) {
		a.createNewFile(fpath)
	} else {
		a.appendToFile(fpath)
	}
}

func (a *App) createNewFile(fpath string) {
	file := openFileCreate(fpath)
	defer closeFile(file)

	sb := &strings.Builder{}

	a.addBlockPackage(sb)
	a.addBlockObject(sb)

	writeToFile(file, sb.String())
}

func (a *App) appendToFile(fpath string) {
	logging.Infof("note: file '%s' already exists\n", a.config.Filename)

	file := openFileRW(fpath)
	defer closeFile(file)

	contents := readFromFile(file)
	sb := &strings.Builder{}

	if !strings.Contains(contents, "package "+a.config.PackageName) {
		a.addBlockPackage(sb)
	} else {
		logging.Infof("ignored package '%s'\n", a.config.PackageName)
	}

	if !strings.Contains(contents, "type "+a.config.ObjectName) {
		a.addBlockObject(sb)
	} else {
		logging.Infof("ignored %s '%s'\n", a.config.ObjectType, a.config.ObjectName)
	}

	writeToFile(file, sb.String())
}
