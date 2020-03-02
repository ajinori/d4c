package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/ajinori/d4c/config"
	"github.com/ajinori/d4c/plugins"
)

var (
	c = config.Default
)

func init() {
	viper.SetConfigName("d4c")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	for _, route := range c.Routes {
		switch route.Method {
		case http.MethodGet:
			e.GET(route.Path, func(c echo.Context) error { return plugins.Run(c, route) })
		}
	}

	if err := http.ListenAndServe(":8080", e); err != nil {
		panic(err)
	}
}
