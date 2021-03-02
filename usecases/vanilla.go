package usecases

import (
	"go.uber.org/zap"

	fizzBuzzModule "fizzbuzz/modules/fizzbuzz"
	statsModule "fizzbuzz/modules/stats"
)

type Vanilla struct {
	fizzBuzz fizzBuzzModule.FizzBuzz
	stats    statsModule.Stats
	logger   *zap.Logger
}

func NewUsecases(fizzBuzz fizzBuzzModule.FizzBuzz, stats statsModule.Stats, logger *zap.Logger) Usecases {
	return &Vanilla{
		fizzBuzz: fizzBuzz,
		stats:    stats,
		logger:   logger,
	}
}
