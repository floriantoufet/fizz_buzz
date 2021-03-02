package usecases

import (
	"errors"

	"go.uber.org/zap"

	"fizzbuzz/domains"
	fizzBuzzModule "fizzbuzz/modules/fizzbuzz"
	statsModule "fizzbuzz/modules/stats"
)

var (
	// ErrInvalidLimit is thrown when limit is negative
	ErrInvalidLimit = errors.New("invalid limit")

	// ErrInvalidModulo is thrown when one of modulus is negative
	ErrInvalidModulo = errors.New("invalid modulo")

	// ErrUnexpected is thrown when something unexpected happened
	ErrUnexpected = errors.New("unexpected error")
)

type Usecases interface {
	// FizzBuzz returns a list of strings with numbers from 1 to limit,
	// where: all multiples of request.FizzModulo are replaced by request.FizzString,
	// all multiples of request.BuzzModulo are replaced by request.BuzzString,
	// all multiples of request.FizzModulo and request.BuzzModulo are replaced by request.String1request.BuzzString
	// returns ErrInvalidLimit if limit is negative
	// returns ErrInvalidModulo if one of modulus is negative
	FizzBuzz(request domains.FizzBuzz) (string, error)

	// RetrieveStats returns the FizzBuzz corresponding to the most used request,
	// as well as the number of hits for this request
	RetrieveStats() ([]domains.FizzBuzz, uint)
}

// Vanilla is default implementation of Usecases interface
type Vanilla struct {
	fizzBuzz fizzBuzzModule.FizzBuzz
	stats    statsModule.Stats
	logger   *zap.Logger
}

// NewUsecases is constructor of Vanilla
func NewUsecases(fizzBuzz fizzBuzzModule.FizzBuzz, stats statsModule.Stats, logger *zap.Logger) Usecases {
	return &Vanilla{
		fizzBuzz: fizzBuzz,
		stats:    stats,
		logger:   logger.Named("Usecases"),
	}
}
