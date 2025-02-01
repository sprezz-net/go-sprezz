package handler

import (
	"net/http"

	"github.com/ankorstore/yokai/config"
	"github.com/ankorstore/yokai/log"
	"github.com/labstack/echo/v4"
)

// WebFingerResponse.
type WebFingerResponse struct {
	Subject    string            `json:"subject"`
	Aliases    []string          `json:"aliases,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Links      []WebFingerLink   `json:"links,omitempty"`
}

// WebFingerLink.
type WebFingerLink struct {
	Rel        string            `json:"rel"`
	Href       string            `json:"href,omitempty"`
	Type       string            `json:"type,omitempty"`
	Titles     map[string]string `json:"titles,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

// WebFingerHandler is an HTTP handler for webfinger requests.
type WebFingerHandler struct {
	config *config.Config
	log    *log.Logger
}

// NewWebFingerHandler returns a new [WebFingerHandler].
func NewWebFingerHandler(config *config.Config, log *log.Logger) *WebFingerHandler {
	return &WebFingerHandler{
		config: config,
		log:    log,
	}
}

// Handle handles HTTP requests.
func (h *WebFingerHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		rsc := c.QueryParam("resource")
		// TODO
		// Implement real resource retrieval
		// Filter for rel params on returned links
		wfr := &WebFingerResponse{
			Subject: rsc,
		}

		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		c.Response().Header().Set(echo.HeaderContentType, "application/jrd+json")

		return c.JSONPretty(http.StatusOK, wfr, "  ")
	}
}
