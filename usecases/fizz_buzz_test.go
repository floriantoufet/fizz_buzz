package usecases_test

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"

	"fizzbuzz/domains"
	"fizzbuzz/modules/fizzbuzz"
	"fizzbuzz/usecases"
)

func TestVanilla_FizzBuzz(t *testing.T) {
	Convey("When I try to get fizzBuzz", t, func() {
		mocks := []*mock.Mock{
			fizzBuzzMock.GetMock(),
			statsMock.GetMock(),
		}

		uc := newUsecasesForUnitTests()

		request := domains.FizzBuzz{
			FizzModulo: 3,
			BuzzModulo: 5,
			Limit:      15,
			FizzString: "fizz",
			BuzzString: "buzz",
		}

		Convey("should success", func() {
			expectedFizzBuzzString := "fizzBuzz"

			fizzBuzzMock.
				On("FizzBuzz", request).
				Return(expectedFizzBuzzString, nil)

			statsMock.On("RecordFizzBuzzRequest", request)

			result, err := uc.FizzBuzz(request)

			So(err, ShouldBeNil)
			So(result, ShouldEqual, expectedFizzBuzzString)
		})

		Convey("should fail", func() {
			var expectedErr error

			Convey("if invalid modulo", func() {
				fizzBuzzMock.
					On("FizzBuzz", request).
					Return("", fizzbuzz.ErrInvalidModulo)

				expectedErr = usecases.ErrInvalidModulo
			})
			Convey("if invalid limit", func() {
				fizzBuzzMock.
					On("FizzBuzz", request).
					Return("", fizzbuzz.ErrInvalidLimit)

				expectedErr = usecases.ErrInvalidLimit
			})
			Convey("if get an unexpected error from fizzBuzz", func() {
				fizzBuzzMock.
					On("FizzBuzz", request).
					Return("", errors.New("ooops"))

				expectedErr = usecases.ErrUnexpected
			})

			result, err := uc.FizzBuzz(request)

			So(result, ShouldEqual, "")
			So(err, ShouldBeError, expectedErr)
		})

		resetMocks(mocks...)
	})
}
