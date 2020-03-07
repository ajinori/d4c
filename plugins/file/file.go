package file

import (
	"github.com/ajinori/d4c/config"
)

type Context interface {
	File(string) error
}

func Run(c Context, route config.Route) error {
	return c.File(route.Plugin.Filename)
}
