// +build wireinject
package advanced

import "github.com/google/wire"

type Fooer interface {
	Foo() string
}

type Bar string

func (b *Bar) Foo() string {
	return string(*b)
}

func ProvideBar() *Bar {
	b := new(Bar)
	*b = "Hello World!"
	return b
}

type FooerPlus interface {
	Fooer
	Bar() string
}

var BarFooer = wire.NewSet(
	ProvideBar,
	wire.Bind(new(Fooer), new(Bar)))

type BarPlus string

func (bp BarPlus) Foo() string { return string(bp) + "in foo" }
func (bp BarPlus) Bar() string { return string(bp) + "in bar" }

func ProvideFooerPlus() FooerPlus {
	return BarPlus("I am barplus")
}

var FooerPlusAsFooer = wire.NewSet(
	ProvideFooerPlus,
	wire.Bind(new(Fooer), *new(FooerPlus)))

func initializeBarFooer() (*Bar, error) {
	wire.Build(BarFooer)
	return nil, nil
}

func initializeBarFooerPlus() (FooerPlus, error) {
	wire.Build(FooerPlusAsFooer)
	return nil, nil
}
