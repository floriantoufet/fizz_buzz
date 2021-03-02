package inject

import (
	"go.uber.org/fx"

	"fizzbuzz/usecases"
)

var UseCases = fx.Provide(
	usecases.NewUsecases,
)
