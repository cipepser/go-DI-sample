// +build wireinject
package cleanupFunctions

import (
	"os"

	"github.com/google/wire"
)

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

func injectFile(path string) (*os.File, func(), error) {
	wire.Build(ProvideLogger, provideFile, ProvidePath)
	return nil, nil, nil
}
