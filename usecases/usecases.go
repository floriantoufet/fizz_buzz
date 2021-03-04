package usecases

import (
	"errors"

	"go.uber.org/zap"

	"github.com/floriantoufet/fizzbuzz/domains"
	fizzBuzzModule "github.com/floriantoufet/fizzbuzz/modules/fizzbuzz"
	statsModule "github.com/floriantoufet/fizzbuzz/modules/stats"
)

const (
	MaxLimitAllowed        = 1000
	MaxStringLengthAllowed = 100
)

var (
	// ErrInvalidLimit is thrown when limit is negative or zero
	ErrInvalidLimit = errors.New("invalid limit")

	// ErrMaxAllowedLimitExceed is thrown when limit is higher than max allowed
	ErrMaxAllowedLimitExceed = errors.New("max allowed limit exceed")

	// ErrMissingLimit is thrown when limit is missing
	ErrMissingLimit = errors.New("missing limit")

	// ErrInvalidFizzModulo is thrown when fizz modulo is zero or negative
	ErrInvalidFizzModulo = errors.New("invalid fizz modulo")

	// ErrMissingFizzModulo is thrown when fizz modulo is missing
	ErrMissingFizzModulo = errors.New("missing fizz modulo")

	// ErrInvalidBuzzModulo is thrown when buzz modulo is zero or negative
	ErrInvalidBuzzModulo = errors.New("invalid buzz modulo")

	// ErrMissingBuzzModulo is thrown when buzz modulo is missing
	ErrMissingBuzzModulo = errors.New("missing buzz modulo")

	// ErrMissingFizzString is thrown when fizz string is missing or empty
	ErrMissingFizzString = errors.New("missing fizz string")

	// ErrMissingBuzzString is thrown when buzz string is missing or empty
	ErrMissingBuzzString = errors.New("missing buzz string")

	// ErrTooLongBuzzString is thrown when buzz string length is greater than
	// max string length allowed
	ErrTooLongBuzzString = errors.New("too long buzz string")

	// ErrTooLongFizzString is thrown when fizz string length is greater than
	// max string length allowed
	ErrTooLongFizzString = errors.New("too long fizz string")

	// ErrUnexpected is thrown when something unexpected happened
	ErrUnexpected = errors.New("unexpected error")
)

type Usecases interface {
	// FizzBuzz returns a list of strings with numbers from 1 to limit where:
	// all multiples of request.FizzModulo are replaced by request.FizzString,
	// all multiples of request.BuzzModulo are replaced by request.BuzzString,
	// all multiples of request.FizzModulo and request.BuzzModulo are replaced by FizzStringBuzzString
	// or error if one of given parameters are invalid
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
