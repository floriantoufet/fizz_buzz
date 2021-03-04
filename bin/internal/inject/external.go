package inject

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/floriantoufet/fizzbuzz/bin/internal"
	"github.com/floriantoufet/fizzbuzz/modules/config"
)

var External = fx.Provide(
	// Config
	config.NewConfig,

	// Zap logger
	func(conf *config.Config) (logger *zap.Logger, err error) {
		switch conf.Logger.Env {
		case config.ProdEnv:
			logger, err = zap.NewProduction()
		case config.DevEnv:
			logger, err = zap.NewDevelopment()
		default:
			// By default testing logger
			logger = zap.NewExample()
		}

		// Replace default Zap logger singleton with new instance
		zap.ReplaceGlobals(logger)
		return logger, nil
	},

	// FX Server
	internal.NewFXServer,
)
