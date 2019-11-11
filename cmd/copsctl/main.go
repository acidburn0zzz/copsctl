package main

//go:generate esc -o ../../pkg/resources/static.go -pkg resources -include=\\*.yaml ../..

import (
	"os"

	"github.com/conplementAG/copsctl/pkg/common/logging"
)

func main() {
	defer logging.Dispose()
	defer errorhandler()

	Execute()
}

func errorhandler() {
	// as this is a CLI tool, and not a library with an API, panic is used for most errors that occur,
	// since they are unrecoverable and need some user intervention (or they are genuine panic programming
	// errors)
	if r := recover(); r != nil {
		logging.Errorf("copsctl --- error occured: %+v\n", r)
		os.Exit(1)
	}
}
