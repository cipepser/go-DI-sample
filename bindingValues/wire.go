// +build wireinject
package bindingValues

import (
	"io"
	"os"

	"github.com/google/wire"
)

type Foo struct {
	X int
}

func injectFoo() Foo {
	wire.Build(wire.Value(Foo{X: 42}))
	return Foo{}
}

func injectReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}
