package text

import (
	"net/http"

	"github.com/ajinori/d4c/config"
)

type Context interface {
	String(int, string) error
}

func Run(c Context, route config.Route) error {
	code := http.StatusOK
	if route.Code > 0 {
		code = route.Code
	}

	return c.String(code, route.Plugin.Message)
}
