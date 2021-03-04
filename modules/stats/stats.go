package stats

import (
	"github.com/floriantoufet/fizzbuzz/domains"
)

type Stats interface {
	// RetrieveMostFrequentFizzBuzzRequest returns the FizzBuzz corresponding to the most used request,
	// as well as the number of hits for this request
	RetrieveMostFrequentFizzBuzzRequest() ([]domains.FizzBuzz, uint)

	// RecordFizzBuzzRequest records given request to called request
	RecordFizzBuzzRequest(request domains.FizzBuzz)
}
