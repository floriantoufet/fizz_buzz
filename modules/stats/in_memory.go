package stats

import (
	"github.com/floriantoufet/fizzbuzz/domains"
)

type InMemory struct {
	maxCalled        uint
	fizzBuzzRequests map[domains.FizzBuzz]uint
}

func NewStats() Stats {
	return &InMemory{
		fizzBuzzRequests: map[domains.FizzBuzz]uint{},
	}
}

// RecordFizzBuzzRequest implements Stats interface
func (stats *InMemory) RecordFizzBuzzRequest(parameters domains.FizzBuzz) {
	stats.fizzBuzzRequests[parameters] ++
	if stats.fizzBuzzRequests[parameters] > stats.maxCalled {
		stats.maxCalled = stats.fizzBuzzRequests[parameters]
	}
}

// RetrieveMostFrequentFizzBuzzRequest implements Stats interface
func (stats InMemory) RetrieveMostFrequentFizzBuzzRequest() (requests []domains.FizzBuzz, count uint) {
	count = stats.maxCalled

	for r, c := range stats.fizzBuzzRequests {
		if c == stats.maxCalled {
			requests = append(requests, r)
		}
	}

	return
}

// ResetStats implements Stats interface
func (stats *InMemory) ResetStats() {
	*stats = InMemory{
		fizzBuzzRequests: map[domains.FizzBuzz]uint{},
	}
}
