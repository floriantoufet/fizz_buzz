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

func (v Vanilla) FizzBuzz(request domains.FizzBuzz) (result string, err error) {
	// Check if Limit is strictly positive
	if request.Limit <= 0 {
		return "", ErrInvalidLimit
	}

	// Negatives modulus are forbidden
	if request.BuzzModulo < 0 || request.FizzModulo < 0 {
		return "", ErrInvalidModulo
	}

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

	result = strings.TrimRight(result, ",")

	return
}
