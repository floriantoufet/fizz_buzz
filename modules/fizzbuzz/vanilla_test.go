package fizzbuzz_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/floriantoufet/fizzbuzz/domains"
	"github.com/floriantoufet/fizzbuzz/modules/fizzbuzz"
)

func TestVanilla_FizzBuzz(t *testing.T) {
	Convey("When I try to get FizzBuzz string", t, func() {
		v := fizzbuzz.NewFizzBuzz()

		Convey("should success with positive modulus", func() {
			request := domains.FizzBuzz{FizzModulo: 3, BuzzModulo: 5, Limit: 16, FizzString: "fizz", BuzzString: "buzz"}

			result, err := v.FizzBuzz(request)

			So(err, ShouldBeNil)
			So(result, ShouldEqual, "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16")
		})

		Convey("should fail", func() {
			var request domains.FizzBuzz
			var expectedErr error

			Convey("with negative modulus", func() {
				request = domains.FizzBuzz{FizzModulo: -3, BuzzModulo: -5, Limit: 16, FizzString: "fizz", BuzzString: "buzz"}
				expectedErr = fizzbuzz.ErrInvalidModulo
			})

			Convey("with invalid limit", func() {
				request = domains.FizzBuzz{FizzModulo: 3, BuzzModulo: 5, Limit: -5, FizzString: "fizz", BuzzString: "buzz"}
				expectedErr = fizzbuzz.ErrInvalidLimit
			})

			result, err := v.FizzBuzz(request)

			So(err, ShouldBeError, expectedErr)
			So(result, ShouldEqual, "")
		})

	})
}
