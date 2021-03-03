package usecases

import (
	"errors"

	"go.uber.org/zap"

	"fizzbuzz/domains"
	fizzBuzzModule "fizzbuzz/modules/fizzbuzz"
)

// FizzBuzz implements Usecases interface
func (uc Vanilla) FizzBuzz(fizzModulo, buzzModulo, limit *int, fizzString, buzzString *string) (string, error) {
	logger := uc.logger.Named("FizzBuzz")

	if fizzModulo == nil || buzzModulo == nil {
		return "", ErrInvalidModulo
	}
	if limit == nil {
		return "", ErrInvalidLimit
	}
	if fizzString == nil || buzzString == nil {
		return "", ErrUnexpected
	}
	request := domains.FizzBuzz{}

	// Get fizzBuzz string
	result, err := uc.fizzBuzz.FizzBuzz(request)
	if err != nil {
		switch {
		case errors.Is(err, fizzBuzzModule.ErrInvalidModulo):
			logger.Debug("Invalid modulo")
			return "", ErrInvalidModulo
		case errors.Is(err, fizzBuzzModule.ErrInvalidLimit):
			logger.Debug("Invalid limit")
			return "", ErrInvalidLimit
		default:
			logger.Error("Unable to get fizzBuzz query", zap.Error(err))
			return "", ErrUnexpected
		}
	}

	// Record request
	uc.stats.RecordFizzBuzzRequest(request)

	logger.Debug("Success")

	return result, nil
}
