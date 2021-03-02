package inject

import (
	"go.uber.org/fx"

	"fiz_buz/usecases"
)

var UseCases = fx.Provide(
	usecases.NewUsecases,
)
