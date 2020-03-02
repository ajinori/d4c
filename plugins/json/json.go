package json

import (
	"encoding/json"
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

	var data interface{}
	if err := json.Unmarshal([]byte(route.Plugin.Content), &data); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(code, data)
}
