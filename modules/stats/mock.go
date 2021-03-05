package stats

import (
	"github.com/stretchr/testify/mock"

	"github.com/floriantoufet/fizzbuzz/domains"
)

var _ Stats = &Mock{}

type Mock struct {
	mock.Mock
}

func (m *Mock) GetMock() *mock.Mock {
	return &m.Mock
}

func (m *Mock) RetrieveMostFrequentFizzBuzzRequest() ([]domains.FizzBuzz, uint) {
	return m.Called().Get(0).([]domains.FizzBuzz), m.Called().Get(1).(uint)
}

func (m *Mock) RecordFizzBuzzRequest(parameters domains.FizzBuzz) {
	_ = m.Called(parameters)
}

func (m *Mock) ResetStats() {
	m.Called()
}
