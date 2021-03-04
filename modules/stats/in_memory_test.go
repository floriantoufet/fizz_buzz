package stats_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/floriantoufet/fizzbuzz/domains"
	"github.com/floriantoufet/fizzbuzz/modules/stats"
)

func Test_InMemory(t *testing.T) {
	Convey("When I try to get most frequent fizz buzz should success", t, func() {
		s := stats.NewStats()

		request1 := domains.FizzBuzz{FizzModulo: 1, BuzzModulo: 4, Limit: 20, FizzString: "fizz", BuzzString: "buzz"}
		request2 := domains.FizzBuzz{FizzModulo: 5, BuzzModulo: 2, Limit: 25, FizzString: "buzz", BuzzString: "leclair"}
		request3 := domains.FizzBuzz{FizzModulo: 6, BuzzModulo: 2, Limit: 25, FizzString: "buzz", BuzzString: "aldrin"}

		var expectedRequests []domains.FizzBuzz
		var expectedCount uint

		Convey("with one request most called", func() {
			s.RecordFizzBuzzRequest(request1)
			s.RecordFizzBuzzRequest(request1)
			s.RecordFizzBuzzRequest(request1)
			s.RecordFizzBuzzRequest(request2)
			s.RecordFizzBuzzRequest(request3)
			s.RecordFizzBuzzRequest(request3)

			expectedRequests = []domains.FizzBuzz{request1}
			expectedCount = 3

		})

		Convey("with multiple most called requests", func() {
			s.RecordFizzBuzzRequest(request1)
			s.RecordFizzBuzzRequest(request1)
			s.RecordFizzBuzzRequest(request1)
			s.RecordFizzBuzzRequest(request3)
			s.RecordFizzBuzzRequest(request3)
			s.RecordFizzBuzzRequest(request3)
			s.RecordFizzBuzzRequest(request2)
			s.RecordFizzBuzzRequest(request2)

			expectedRequests = []domains.FizzBuzz{request1, request3}
			expectedCount = 3
		})

		Convey("with multiple most called requests and reset", func() {
			s.RecordFizzBuzzRequest(request1)
			s.RecordFizzBuzzRequest(request1)
			s.RecordFizzBuzzRequest(request1)
			s.RecordFizzBuzzRequest(request3)
			s.RecordFizzBuzzRequest(request3)
			s.RecordFizzBuzzRequest(request3)
			s.RecordFizzBuzzRequest(request2)
			s.RecordFizzBuzzRequest(request2)
			s.ResetStats()

			expectedRequests = []domains.FizzBuzz{}
			expectedCount = 0
		})

		requests, count := s.RetrieveMostFrequentFizzBuzzRequest()

		So(len(requests), ShouldEqual, len(expectedRequests))
		found := len(expectedRequests)

		for _, r := range requests {
			for _, expected := range expectedRequests {
				if expected == r {
					found--
				}
			}
		}

		So(found, ShouldEqual, 0)
		So(count, ShouldEqual, expectedCount)
	})
}
