package domains_test

import (
	"encoding/json"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/floriantoufet/fizzbuzz/domains"
)

func TestError(t *testing.T) {
	Convey("When I try to get errors", t, func() {
		errs := domains.Errors{}

		err1 := errors.New("error1")
		err2 := errors.New("error2")

		So(errs.IsEmpty(), ShouldBeTrue)

		errs.Add(err1)
		errs.Add(err2)

		So(errs.IsEmpty(), ShouldBeFalse)

		So(errs.Contains(err1), ShouldBeTrue)
		So(errs.Contains(err2), ShouldBeTrue)

		So(errs.Error(), ShouldEqual, "error1, error2")

		raw, err := json.Marshal(errs)
		So(err, ShouldBeNil)
		So(string(raw), ShouldEqual, `["error1","error2"]`)
	})
}
