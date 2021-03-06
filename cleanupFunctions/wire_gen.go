// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cleanupFunctions

import (
	"os"
)

// Injectors from wire.go:

func injectFile(path string) (*os.File, func(), error) {
	logger := ProvideLogger()
	cleanupFunctionsPath := ProvidePath(path)
	file, cleanup, err := provideFile(logger, cleanupFunctionsPath)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		cleanup()
	}, nil
}

// wire.go:

type Logger struct{}

func (l Logger) Log(err error) {}

func ProvideLogger() Logger {
	return Logger{}
}

type Path string

func ProvidePath(path string) Path {
	return Path(path)
}

func provideFile(log Logger, path Path) (*os.File, func(), error) {
	f, err := os.Open(string(path))
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err := f.Close(); err != nil {
			log.Log(err)
		}
	}
	return f, cleanup, nil
}
