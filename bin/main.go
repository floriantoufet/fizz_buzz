package main

import (
	"go.uber.org/fx"

	"fiz_buz/bin/internal"
	"fiz_buz/bin/internal/inject"
)

func main() {
	app := fx.New(
		// fx.NopLogger, // remove for debug
		inject.External,
		inject.Modules,
		inject.Transports,
		inject.UseCases,
		// fx.Invoke(validation.ValidateConfig),
		fx.Invoke(internal.Dependencies),
		fx.Invoke(internal.HTTPServer),
	)

	internal.Start(app)
	<-app.Done()
	internal.Shutdown(app)
}
