package json

import (
	"net/http"

	"github.com/ajinori/d4c/config"
)

type Context interface {
	JSON(int, interface{}) error
}

func Run(c Context, route config.Route) error {
	code := http.StatusOK
	if route.Code > 0 {
		code = route.Code
	}

	return c.JSON(code, route.Plugin.Data)
}
