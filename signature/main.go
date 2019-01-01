// +build wireinject
package signature

import (
	"context"

	"github.com/google/wire"
)

func initializeBaz(ctx context.Context) (Baz, error) {
	wire.Build(SuperSet)
	return Baz{}, nil
}
