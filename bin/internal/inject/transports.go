package inject

import (
	"go.uber.org/fx"

	"fizzbuzz/transport/http"
	"fizzbuzz/transport/http/endpoints"
)

var Transports = fx.Options(
	fx.Provide(
		// Force annotation http instance to "http" to provide multiple transports
		// implementation
		fx.Annotated{Name: "http", Target: http.NewHTTP},

		// HTTP Config
		func() *http.Config {
			return &http.Config{
				Host: "localhost",
				Port: 8080,
			}
		},
	),

	// Endpoints
	fx.Provide(endpoints.NewEndpoints),
)
