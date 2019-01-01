// +build wireinject
package structProviders

import "github.com/google/wire"

type Foo int
type Bar int

func ProvideFoo() Foo { return Foo(10) }
func ProvideBar() Bar { return Bar(20) }

type FooBar struct {
	Foo Foo
	Bar Bar
}

var Set = wire.NewSet(
	ProvideFoo,
	ProvideBar,
	FooBar{})

func initializeFooBar() (FooBar, error) {
	wire.Build(Set)
	return FooBar{}, nil
}
