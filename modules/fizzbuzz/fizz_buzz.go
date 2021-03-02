package fizzbuzz

import (
	"errors"

	"fizzbuzz/domains"
)

var (
	ErrInvalidModulo = errors.New("invalid modulo")
	ErrInvalidLimit  = errors.New("invalid limit")
)

type FizzBuzz interface {
	// FizzBuzz returns a list of strings with numbers from 1 to limit,
	// where: all multiples of request.FizzModulo are replaced by request.FizzString,
	// all multiples of request.BuzzModulo are replaced by request.BuzzString,
	// all multiples of request.FizzModulo and request.BuzzModulo are replaced by request.String1request.BuzzString
	FizzBuzz(request domains.FizzBuzz) (string, error)
}
