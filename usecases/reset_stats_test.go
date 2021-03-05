package usecases_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVanilla_ResetStats(t *testing.T) {
	Convey("When I try to reset stats should success", t, func() {
		uc := newUsecasesForUnitTests()

		statsMock.On("ResetStats")

		uc.ResetStats()
		
		So(statsMock.AssertCalled(t, "ResetStats"), ShouldBeTrue)

		resetMock(statsMock.GetMock())
	})
}
