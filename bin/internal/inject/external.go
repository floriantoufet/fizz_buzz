package inject

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/floriantoufet/fizzbuzz/bin/internal"
)

var External = fx.Provide(
	// Zap logger
	func() (*zap.Logger, error) {
		logger := zap.NewExample()

		// Replace default Zap logger singleton with new instance
		zap.ReplaceGlobals(logger)
		return logger, nil
	},

	// FX Server
	internal.NewFXServer,
)
