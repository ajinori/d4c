package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/ajinori/d4c/config"
	"github.com/ajinori/d4c/modules/text"
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
	fmt.Println(c)
	e := echo.New()

	for _, route := range c.Routes {
		switch route.Method {
		case http.MethodGet:
			e.GET(route.Path, func(c echo.Context) error { return text.Module(c, route.Text.Message) })
		}
	}

	e.Start(":8080")
}
