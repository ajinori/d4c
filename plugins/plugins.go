package plugins

import (
	"github.com/labstack/echo"

	"github.com/ajinori/d4c/config"
	"github.com/ajinori/d4c/plugins/html"
	"github.com/ajinori/d4c/plugins/json"
	"github.com/ajinori/d4c/plugins/redirect"
	"github.com/ajinori/d4c/plugins/text"
)

func Run(route config.Route) func(echo.Context) error {
	return func(c echo.Context) error {
		switch route.Plugin.Name {
		case "text":
			return text.Run(c, route)
		case "html":
			return html.Run(c, route)
		case "json":
			return json.Run(c, route)
		case "redirect":
			return redirect.Run(c, route)
		}

		return nil
	}
}
