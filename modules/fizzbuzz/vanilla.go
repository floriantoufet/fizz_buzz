package fizzbuzz

import (
	"strconv"
	"strings"

	"fizzbuzz/domains"
)

type Vanilla struct {
}

func NewFizzBuzz() FizzBuzz {
	return &Vanilla{}
}

// FizzBuzz implements FizzBuzz interface
func (v Vanilla) FizzBuzz(request domains.FizzBuzz) (string, error) {
	// Check if Limit is strictly positive
	if request.Limit <= 0 {
		return "", ErrInvalidLimit
	}

	// Zero or negatives modulus are forbidden
	if request.BuzzModulo <= 0 || request.FizzModulo <= 0 {
		return "", ErrInvalidModulo
	}

	// Build fizzBuzz string
	var result string
	for n := 1; n <= request.Limit; n++ {
		current := ""

		if n%request.FizzModulo == 0 {
			current += request.FizzString
		}

		if n%request.BuzzModulo == 0 {
			current += request.BuzzString
		}

		if current == "" {
			current += strconv.Itoa(n)
		}

		result += current + ","
	}

	return strings.TrimRight(result, ","), nil
}
