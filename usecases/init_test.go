package usecases_test

import (
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"

	"github.com/floriantoufet/fizzbuzz/modules/fizzbuzz"
	"github.com/floriantoufet/fizzbuzz/modules/stats"
	"github.com/floriantoufet/fizzbuzz/usecases"
)

var (
	fizzBuzzMock = &fizzbuzz.Mock{}
	statsMock    = &stats.Mock{}
)

func newUsecasesForUnitTests() usecases.Usecases {
	return usecases.NewUsecases(fizzBuzzMock, statsMock, zap.NewNop())
}

// resetMocks reset assertions for provided mock list
func resetMocks(mocks ...*mock.Mock) {
	for _, m := range mocks {
		resetMock(m)
	}
}

// resetMock reset assertion for provided mock
func resetMock(m *mock.Mock) {
	m.ExpectedCalls = []*mock.Call{}
	m.Calls = []mock.Call{}
}
