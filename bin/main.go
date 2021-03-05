package main

import (
	"go.uber.org/fx"

	"github.com/floriantoufet/fizzbuzz/bin/internal"
	"github.com/floriantoufet/fizzbuzz/bin/internal/inject"
)

func main() {
	app := fx.New(
		fx.NopLogger, // remove for debug
		inject.External,
		inject.Modules,
		inject.Transports,
		inject.UseCases,
		fx.Invoke(internal.Dependencies),
		fx.Invoke(internal.HTTPServer),
	)

	internal.Start(app)
	<-app.Done()
	internal.Shutdown(app)
}
