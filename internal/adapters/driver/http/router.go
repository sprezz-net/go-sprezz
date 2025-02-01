package http

import (
	"github.com/ankorstore/yokai/fxhttpserver"
	"github.com/sprezz-net/go-sprezz/internal/adapters/driver/http/handler"
	"go.uber.org/fx"
)

// Router is used to register the application HTTP middlewares and handlers.
func Router() fx.Option {
	return fx.Options(
		fxhttpserver.AsHandler("GET", "/", handler.NewRootHandler),
		fxhttpserver.AsHandler("GET", "/.well-known/webfinger", handler.NewWebFingerHandler),
	)
}
