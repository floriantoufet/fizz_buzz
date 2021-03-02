package stats

import (
	"fizzbuzz/domains"
)

type Stats interface {
	RetrieveMostFrequentFizzBuzzRequest() ([]domains.FizzBuzz, uint)
	RecordFizzBuzzRequest(parameters domains.FizzBuzz)
}
