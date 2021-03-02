package fizzbuzz

import (
	"github.com/stretchr/testify/mock"

	"fizzbuzz/domains"
)

var _ FizzBuzz = &Mock{}

type Mock struct {
	mock.Mock
}

func (m *Mock) FizzBuzz(request domains.FizzBuzz) (string, error) {
	args := m.Called(request)

	return args.String(0), args.Error(1)
}
