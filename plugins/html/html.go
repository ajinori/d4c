package html

import (
	"net/http"

	"github.com/ajinori/d4c/config"
)

type Context interface {
	HTML(int, string) error
}

func Run(c Context, route config.Route) error {
	code := http.StatusOK
	if route.Code > 0 {
		code = route.Code
	}

	return c.HTML(code, route.Plugin.Content)
}
