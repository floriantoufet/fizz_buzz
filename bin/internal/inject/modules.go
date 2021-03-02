package inject

import (
	"go.uber.org/fx"

	"fizzbuzz/modules/fizzbuzz"
	"fizzbuzz/modules/stats"
)

var Modules = fx.Provide(
	fizzbuzz.NewFizzBuzz,
	stats.NewStats,
)
