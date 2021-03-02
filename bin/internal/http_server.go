package internal

import (
	"go.uber.org/fx"

	"fizzbuzz/transport"
)

// HTTPServerParams is the input parameter struct for the modules that contains its dependencies.
type HTTPServerParams struct {
	fx.In
	Server    *FXServer
	Transport transport.Server `name:"http"`
}

// HTTPServer registers the routes for the server and starts the server on app start.
// uber-fx does not supports pointers while gocritics expect a pointer for heavy struct
// nolint:gocritic
func HTTPServer(p HTTPServerParams) {
	p.Server.Run("HTTPServer", p.Transport)
}
