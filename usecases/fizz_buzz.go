package usecases

import (
	"unicode/utf8"

	"go.uber.org/zap"

	"github.com/floriantoufet/fizzbuzz/domains"
)

// FizzBuzz implements Usecases interface
func (uc Vanilla) FizzBuzz(fizzModulo, buzzModulo, limit *int, fizzString, buzzString *string) (string, error) {
	logger := uc.logger.Named("FizzBuzz")

	// Build request
	request, requestErr := uc.getRequest(fizzModulo, buzzModulo, limit, fizzString, buzzString)
	if requestErr != nil {
		logger.Debug("Invalid request", zap.Error(requestErr))
		return "", requestErr
	}

	// Get fizzBuzz string
	result, err := uc.fizzBuzz.FizzBuzz(request)
	if err != nil {
		logger.Error("Unable to get fizzBuzz query", zap.Error(err))
		return "", ErrUnexpected
	}

	// Record request
	uc.stats.RecordFizzBuzzRequest(request)

	logger.Debug("Success")

	return result, nil
}

// getRequest returns domains.FizzBuzz based on given parameters
// or a list of errors if one of parameters are invalid or missing
func (Vanilla) getRequest(fizzModulo, buzzModulo, limit *int, fizzString, buzzString *string) (domains.FizzBuzz, domains.Errors) {
	errs := domains.Errors{}

	// Check modulus
	if fizzModulo == nil {
		errs.Add(ErrMissingFizzModulo)
	} else if *fizzModulo <= 0 {
		errs.Add(ErrInvalidFizzModulo)
	}
	
	if buzzModulo == nil {
		errs.Add(ErrMissingBuzzModulo)
	} else if *buzzModulo <= 0 {
		errs.Add(ErrInvalidBuzzModulo)
	}

	// Check limit
	switch {
	case limit == nil:
		errs.Add(ErrMissingLimit)
		break
	case *limit <= 0:
		errs.Add(ErrInvalidLimit)
		break
	case *limit > MaxLimitAllowed:
		errs.Add(ErrMaxAllowedLimitExceed)
	}

	// Check string
	if fizzString == nil {
		errs.Add(ErrMissingFizzString)
	} else if utf8.RuneCountInString(*fizzString) > MaxStringLengthAllowed {
		errs.Add(ErrTooLongFizzString)
	}

	if buzzString == nil {
		errs.Add(ErrMissingBuzzString)
	} else if utf8.RuneCountInString(*buzzString) > MaxStringLengthAllowed {
		errs.Add(ErrTooLongBuzzString)
	}

	if !errs.IsEmpty() {
		return domains.FizzBuzz{}, errs
	}

	return domains.FizzBuzz{
		FizzModulo: *fizzModulo,
		BuzzModulo: *buzzModulo,
		Limit:      *limit,
		FizzString: *fizzString,
		BuzzString: *buzzString,
	}, nil
}
