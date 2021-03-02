package stats

import (
	"github.com/stretchr/testify/mock"

	"fizzbuzz/domains"
)

var _ Stats = &Mock{}

type Mock struct {
	mock.Mock
}

func (m *Mock) RetrieveMostFrequentFizzBuzzRequest() ([]domains.FizzBuzz, uint) {
	return m.Called().Get(0).([]domains.FizzBuzz), m.Called().Get(1).(uint)
}

func (m *Mock) RecordFizzBuzzRequest(parameters domains.FizzBuzz) {
	_ = m.Called(parameters)
}
