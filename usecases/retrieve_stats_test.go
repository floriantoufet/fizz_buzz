package usecases_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"fizzbuzz/domains"
)

func TestVanilla_RetrieveStats(t *testing.T) {
	Convey("When I try to retrieve stats should success", t, func() {
		uc := newUsecasesForUnitTests()

		expectedStats := []domains.FizzBuzz{
			{
				FizzModulo: 3,
				BuzzModulo: 5,
				Limit:      15,
				FizzString: "fizz",
				BuzzString: "buzz",
			},
			{
				FizzModulo: 2,
				BuzzModulo: 7,
				Limit:      37,
				FizzString: "buzz",
				BuzzString: "leclair",
			},
		}
		expectedTotal := uint(5)

		statsMock.
			On("RetrieveMostFrequentFizzBuzzRequest").
			Return(
				expectedStats,
				expectedTotal,
			)

		result, total := uc.RetrieveStats()

		So(result, ShouldResemble, expectedStats)
		So(total, ShouldEqual, expectedTotal)

		resetMock(statsMock.GetMock())
	})
}
