package redirect

import (
	"net/http"
	"net/url"

	"github.com/ajinori/d4c/config"
)

type Context interface {
	String(int, string) error
	Redirect(int, string) error
}

func Run(c Context, route config.Route) error {
	u, err := url.Parse(route.Plugin.URL)
	if err != nil {
		return c.String(http.StatusInternalServerError, "invalid url: "+route.Plugin.URL)
	}
	return c.Redirect(http.StatusMovedPermanently, u.String())
}
