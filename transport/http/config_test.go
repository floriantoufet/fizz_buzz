package http_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"fiz_buz/transport/http"
)

func TestUnit_Config_Check(t *testing.T) {
	Convey("Given a configuration", t, func() {
		config := http.Config{
			Host: "localhost",
			Port: 8080,
		}

		Convey("Should success", func() {
			So(config.Check(), ShouldBeNil)
		})

		Convey("Should fail with an invalid TCP port", func() {
			config.Port = 65536

			So(config.Check(), ShouldBeError, http.ErrInvalidPort)
		})
	})
}
