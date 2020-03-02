package plugins

import (
	"github.com/labstack/echo"

	"github.com/ajinori/d4c/config"
	"github.com/ajinori/d4c/plugins/text"
)

func Run(route config.Route) func(echo.Context) error {
	return func(c echo.Context) error {
		switch route.Plugin.Name {
		case "text":
			return text.Run(c, route)
		}

		return nil
	}
}
