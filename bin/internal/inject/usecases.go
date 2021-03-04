package inject

import (
	"go.uber.org/fx"

	"github.com/floriantoufet/fizzbuzz/usecases"
)

var UseCases = fx.Provide(
	usecases.NewUsecases,
)
