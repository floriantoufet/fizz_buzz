package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"fizzbuzz/transport"
	"fizzbuzz/transport/http/endpoints"
)

type HTTP struct {
	config    *Config
	server    *http.Server
	endpoints *endpoints.Endpoint
}

func NewHTTP(httpConfig *Config, endpoints *endpoints.Endpoint) transport.Server {
	return &HTTP{
		config:    httpConfig,
		endpoints: endpoints,
	}
}

func (transport *HTTP) ListenAndServe() error {
	handler := chi.NewRouter()
	handler.Group(func(r chi.Router) {
		transport.initRoutes(r)
	})

	transport.server = &http.Server{
		Addr:    transport.config.GetAddress(),
		Handler: handler,
	}

	return transport.server.ListenAndServe()
}

func (transport *HTTP) Shutdown() error {
	return transport.server.Shutdown(context.Background())
}

func (transport *HTTP) GetAddress() string {
	return transport.config.GetAddress()
}

func (transport *HTTP) initRoutes(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Get("/ping", transport.endpoints.Ping)
	})
}
