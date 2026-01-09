package mkpkg

import (
	"bytes"
	"errors"
	"os"

	"github.com/mrumyantsev/go-errlib"
	"github.com/mrumyantsev/mkpkg/internal/logging"
)

func makeAllDirs(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		logging.ErrorBadquit(errlib.Wrap(err, "could not make package directory"))
	}
}

func isFileExists(fpath string) bool {
	_, err := os.Stat(fpath)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	if err != nil {
		logging.ErrorBadquit(errlib.Wrap(err, "stat"))
	}

	return true
}

func openFileCreate(fpath string) *os.File {
	return openFile(fpath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY)
}

func openFileRW(fpath string) *os.File {
	return openFile(fpath, os.O_RDWR)
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		logging.Error(errlib.Wrap(err, "could not close file"))
	}
}

func readFromFile(file *os.File) string {
	buf := &bytes.Buffer{}

	_, err := buf.ReadFrom(file)
	if err != nil {
		logging.ErrorBadquit(errlib.Wrap(err, "could not read from file"))
	}

	return buf.String()
}

func writeToFile(file *os.File, contents string) {
	_, err := file.WriteString(contents)
	if err != nil {
		logging.ErrorBadquit(errlib.Wrap(err, "could not write to file"))
	}
}

func openFile(fpath string, flag int) *os.File {
	file, err := os.OpenFile(fpath, flag, 0644)
	if err != nil {
		logging.ErrorBadquit(errlib.Wrap(err, "could not create file"))
	}

	return file
}
