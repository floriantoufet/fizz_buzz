package stats

import (
	"fizzbuzz/domains"
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

func (stats *InMemory) RecordFizzBuzzRequest(parameters domains.FizzBuzz) {
	stats.fizzBuzzRequests[parameters] ++
	if stats.fizzBuzzRequests[parameters] > stats.maxCalled {
		stats.maxCalled = stats.fizzBuzzRequests[parameters]
	}
}

func (stats InMemory) RetrieveMostFrequentFizzBuzzRequest() (requests []domains.FizzBuzz, count uint) {
	count = stats.maxCalled

	for r, c := range stats.fizzBuzzRequests {
		if c == stats.maxCalled {
			requests = append(requests, r)
		}
	}

	return
}
