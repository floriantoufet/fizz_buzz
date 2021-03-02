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
	FizzBuzz(request domains.FizzBuzz) (string, error)
}
