package internal

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/floriantoufet/fizzbuzz/transport"
)

type FXServer struct {
	Lifecycle  fx.Lifecycle
	Shutdowner fx.Shutdowner
	Logger     *zap.Logger
}

func NewFXServer(lifecycle fx.Lifecycle, shutdowner fx.Shutdowner, logger *zap.Logger) *FXServer {
	return &FXServer{Lifecycle: lifecycle, Shutdowner: shutdowner, Logger: logger}
}

func (server FXServer) Run(name string, transport transport.Server) {
	logger := server.Logger.
		Named("Lifecycle").
		With(zap.String("address", transport.GetAddress()))

	server.Lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					if err := transport.ListenAndServe(); err != nil {
						if !ShuttingDown {
							logger.Error("🧨💥 "+name+" closed unexpectedly", zap.Error(err))
						}

						if err = server.Shutdowner.Shutdown(); err != nil {
							logger.Error("🧨💥 Unable to shutdown properly "+name, zap.Error(err))
						}
					}
				}()

				logger.Info("📢 " + name + " started")

				return nil
			},

			OnStop: func(context.Context) error {
				if err := transport.Shutdown(); err != nil {
					return err
				}

				logger.Info("📢 " + name + " closed")

				return nil
			},
		},
	)
}
