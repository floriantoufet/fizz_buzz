package usecases_test

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"

	"fizzbuzz/domains"
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

			result, err := uc.FizzBuzz(
				&request.FizzModulo,
				&request.BuzzModulo,
				&request.Limit,
				&request.FizzString,
				&request.BuzzString,
			)

			So(err, ShouldBeNil)
			So(result, ShouldEqual, expectedFizzBuzzString)
		})

		Convey("should fail", func() {
			var (
				expectedErr                   error
				fizzModulo, buzzModulo, limit *int
				fizzString, buzzString        *string
			)

			Convey("if missing parameters", func() {
				errs := domains.Errors{}
				errs.Add(usecases.ErrMissingFizzModulo)
				errs.Add(usecases.ErrMissingBuzzModulo)
				errs.Add(usecases.ErrMissingLimit)
				errs.Add(usecases.ErrMissingFizzString)
				errs.Add(usecases.ErrMissingBuzzString)

				expectedErr = errs
			})

			Convey("if invalid parameters", func() {
				errs := domains.Errors{}
				errs.Add(usecases.ErrInvalidFizzModulo)
				errs.Add(usecases.ErrInvalidBuzzModulo)
				errs.Add(usecases.ErrInvalidLimit)

				nb := -5
				fizzModulo, buzzModulo, limit = &nb, &nb, &nb
				fizzString, buzzString = &request.FizzString, &request.BuzzString

				expectedErr = errs
			})

			Convey("if limit exceed max allowed", func() {
				errs := domains.Errors{}
				errs.Add(usecases.ErrMaxAllowedLimitExceed)

				nb := 1001
				fizzModulo, buzzModulo, limit = &request.FizzModulo, &request.BuzzModulo, &nb
				fizzString, buzzString = &request.FizzString, &request.BuzzString

				expectedErr = errs
			})

			Convey("if get an un unexpected error from module", func() {
				fizzBuzzMock.On("FizzBuzz", request).Return("", errors.New("ooops"))

				fizzModulo, buzzModulo, limit = &request.FizzModulo, &request.BuzzModulo, &request.Limit
				fizzString, buzzString = &request.FizzString, &request.BuzzString

				expectedErr = usecases.ErrUnexpected
			})

			result, err := uc.FizzBuzz(
				fizzModulo,
				buzzModulo,
				limit,
				fizzString,
				buzzString,
			)

			So(result, ShouldEqual, "")
			So(err, ShouldBeError, expectedErr)
		})

		resetMocks(mocks...)
	})
}
