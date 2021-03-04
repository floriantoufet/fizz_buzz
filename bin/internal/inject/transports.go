package inject

import (
	"go.uber.org/fx"

	"github.com/floriantoufet/fizzbuzz/modules/config"
	"github.com/floriantoufet/fizzbuzz/transport/http"
	"github.com/floriantoufet/fizzbuzz/transport/http/endpoints"
)

var Transports = fx.Options(
	fx.Provide(
		// Force annotation http instance to "http" to provide multiple transports
		// implementation
		fx.Annotated{Name: "http", Target: http.NewHTTP},

		// HTTP Config
		func(conf *config.Config) *http.Config {
			return &http.Config{
				Host: conf.Transports.HTTP.Host,
				Port: conf.Transports.HTTP.Port,
			}
		},
	),

	// Endpoints
	fx.Provide(endpoints.NewEndpoints),
)
