package text

import (
	"net/http"

	"github.com/ajinori/d4c/config"
)

type Context interface {
	String(int, string) error
}

func Run(c Context, route config.Route) error {
	return c.String(http.StatusOK, route.Plugin.Message)
}
