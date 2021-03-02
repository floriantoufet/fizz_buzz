package main

import (
	"go.uber.org/fx"

	"fizzbuzz/bin/internal"
	"fizzbuzz/bin/internal/inject"
)

func main() {
	app := fx.New(
		// TODO remove comment
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
