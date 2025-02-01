package handler

import (
	"net/http"

	"github.com/ankorstore/yokai/config"
	"github.com/ankorstore/yokai/log"
	"github.com/labstack/echo/v4"
)

// RootHandler is an HTTP handler for the main page.
type RootHandler struct {
	config *config.Config
	log    *log.Logger
}

// NewRootHandler returns a new [RootHandler].
func NewRootHandler(config *config.Config, log *log.Logger) *RootHandler {
	return &RootHandler{
		config: config,
		log:    log,
	}
}

// Handle handles HTTP requests.
func (h *RootHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Info().Msg("Root")

		return c.HTML(http.StatusOK, "<html><body><h1>Sprezz</h1></body></html>")
	}
}
