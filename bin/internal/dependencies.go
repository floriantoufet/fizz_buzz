package internal

import (
	"context"
	"errors"
	"runtime"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var ErrShutdownDependencies = errors.New("unable to shutdown at least one dependency")

type DependenciesParams struct {
	fx.In
	Lifecycle fx.Lifecycle

	Logger *zap.Logger
}

func Dependencies(p DependenciesParams) {
	logger := p.Logger.Named("Lifecycle")

	p.Lifecycle.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				hasError := false

				// Flush logger (Must be the last to ensure logging)
				// /!\ Failures to sync stdout and stderr on MacOSX:
				// Users can't do anything about failures to sync stderr and stdout.
				// See https://github.com/uber-go/zap/issues/328 and https://github.com/uber-go/zap/issues/370
				bug := "sync /dev/stderr: inappropriate ioctl for device"
				if err := logger.Sync(); err != nil && !(runtime.GOOS == "darwin" && err.Error() == bug) {
					hasError = true
					logger.Error("🧨💥 Unable to sync logger", zap.Error(err))
				}

				if hasError {
					return ErrShutdownDependencies
				}

				logger.Info("📢 Dependencies closed")

				return nil
			},
		},
	)
}
