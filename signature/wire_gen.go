// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package signature

import (
	"context"
)

// Injectors from main.go:

func initializeBaz(ctx context.Context) (Baz, error) {
	foo := ProvideFoo()
	bar := ProvideBar(foo)
	baz, err := ProvideBaz(ctx, bar)
	if err != nil {
		return Baz{}, err
	}
	return baz, nil
}
