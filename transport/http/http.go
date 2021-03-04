package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	gochiCors "github.com/go-chi/cors"

	"github.com/floriantoufet/fizzbuzz/transport"
	"github.com/floriantoufet/fizzbuzz/transport/http/endpoints"
)

type HTTP struct {
	config    *Config
	server    *http.Server
	endpoints *endpoints.Endpoints
}

func NewHTTP(httpConfig *Config, endpoints *endpoints.Endpoints) transport.Server {
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
	// Init CORS
	r.Use(gochiCors.New(gochiCors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
	}).Handler)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/ping", transport.endpoints.Ping)
		r.Get("/fizz_buzz", transport.endpoints.FizzBuzz)
		r.Get("/stats", transport.endpoints.RetrieveStats)
	})
}
