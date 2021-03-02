package usecases

import (
	"errors"

	"fizzbuzz/domains"
	fizzBuzzModule "fizzbuzz/modules/fizzbuzz"
)

func (v Vanilla) FizzBuzz(request domains.FizzBuzz) (string, error) {
	// Get fizzBuzz string
	result, err := v.fizzBuzz.FizzBuzz(request)
	if err != nil {
		switch {
		case errors.Is(err, fizzBuzzModule.ErrInvalidModulo):
			return "", ErrInvalidModulo
		case errors.Is(err, fizzBuzzModule.ErrInvalidLimit):
			return "", ErrInvalidLimit
		default:
			return "", ErrUnexpected
		}
	}

	// Record request
	v.stats.RecordFizzBuzzRequest(request)

	return result, nil
}
