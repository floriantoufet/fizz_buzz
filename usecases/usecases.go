package usecases

import (
	"errors"

	"fizzbuzz/domains"
)

var (
	ErrInvalidLimit  = errors.New("invalid limit")
	ErrInvalidModulo = errors.New("invalid modulo")
	ErrUnexpected    = errors.New("unexpected error")
)

type Usecases interface {
	// Ping will be used for technical purpose
	Ping() string

	// FizzBuzz returns a list of strings with numbers from 1 to limit,
	// where: all multiples of parameters.FizzModulo are replaced by parameters.FizzString,
	// all multiples of parameters.BuzzModulo are replaced by parameters.BuzzString,
	// all multiples of parameters.FizzModulo and parameters.BuzzModulo are replaced by parameters.String1parameters.BuzzString
	FizzBuzz(parameters domains.FizzBuzz) (string, error)

	// RetrieveStats returns the FizzBuzz corresponding to the most used request,
	// as well as the number of hits for this request
	RetrieveStats() (domains.FizzBuzz, int8)
}
