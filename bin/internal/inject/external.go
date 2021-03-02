package inject

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"fiz_buz/bin/internal"
)

var External = fx.Provide(
	// Zap logger
	func() (*zap.Logger, error) {
		logger, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}

		// Replace default Zap logger singleton with new instance
		zap.ReplaceGlobals(logger)
		return logger, nil
	},

	// FX Server
	internal.NewFXServer,
)
