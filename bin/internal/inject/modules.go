package inject

import (
	"go.uber.org/fx"

	"github.com/floriantoufet/fizzbuzz/modules/fizzbuzz"
	"github.com/floriantoufet/fizzbuzz/modules/stats"
)

var Modules = fx.Provide(
	fizzbuzz.NewFizzBuzz,
	stats.NewStats,
)
