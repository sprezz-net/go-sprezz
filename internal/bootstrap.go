package internal

import (
	"context"
	"fmt"
	"testing"

	"github.com/ankorstore/yokai/fxconfig"
	"github.com/ankorstore/yokai/fxcore"
	"github.com/ankorstore/yokai/fxhttpserver"
	"github.com/sprezz-net/go-sprezz/internal/adapters/driver/http"
	"go.uber.org/fx"
)

func init() {
	RootDir = fxcore.RootDir(1)
}

// RootDir is the application root directory.
var RootDir string

// Bootstrapper can be used to load modules, options, dependencies, routing and bootstraps the application.
var ServerBootstrapper = fxcore.NewBootstrapper().WithOptions(
	// modules registration
	fxhttpserver.FxHttpServerModule,
	http.Module,
	// dependencies registration
	Register(),
)

// Run starts the application, with a provided [context.Context].
func RunServer(ctx context.Context) {
	ServerBootstrapper.WithContext(ctx).RunApp()
}

// RunTest starts the application in test mode, with an optional list of [fx.Option].
func RunServerTest(tb testing.TB, options ...fx.Option) {
	tb.Helper()

	ServerBootstrapper.RunTestApp(
		tb,
		// config lookup
		fxconfig.AsConfigPath(fmt.Sprintf("%s/configs/", RootDir)),
		// test options
		fx.Options(options...),
	)
}
