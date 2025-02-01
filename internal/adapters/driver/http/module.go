package http

import (
	"go.uber.org/fx"
)

// Module exports dependency.
var Module = fx.Options(
	Router(),
)
