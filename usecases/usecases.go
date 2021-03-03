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

	// ErrMissingLimit is thrown when limit is missing
	ErrMissingLimit = errors.New("missing limit")

	// ErrInvalidModulo is thrown when one of modulus is zero or negative
	ErrInvalidModulo = errors.New("invalid modulo")

	// ErrInvalidModulo is thrown when one of modulus is zero or negative
	ErrMissingFizzModulo = errors.New("invalid modulo")

	// ErrMissingString is thrown when something unexpected happened
	ErrMissingString = errors.New("unexpected error")

	// ErrUnexpected is thrown when something unexpected happened
	ErrUnexpected = errors.New("unexpected error")
)

type Usecases interface {
	// FizzBuzz returns a list of strings with numbers from 1 to limit where:
	// all multiples of request.FizzModulo are replaced by request.FizzString,
	// all multiples of request.BuzzModulo are replaced by request.BuzzString,
	// all multiples of request.FizzModulo and request.BuzzModulo are replaced by FizzStringBuzzString
	// returns ErrInvalidLimit if limit is negative
	// returns ErrInvalidModulo if one of modulus is zero or negative
	FizzBuzz(fizzModulo, buzzModulo, limit *int, fizzString, buzzString *string) (string, error)

	// RetrieveStats returns the FizzBuzz corresponding to most used requests,
	// as well as the number of hits for those requests
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
