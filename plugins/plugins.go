package plugins

import (
	"github.com/ajinori/d4c/config"
	"github.com/ajinori/d4c/plugins/text"
)

type Context interface {
	String(int, string) error
}

func Run(c Context, route config.Route) error {
	switch route.Plugin.Name {
	case "text":
		return text.Run(c, route)
	}

	return nil
}
