package main

import (
	"github.com/labstack/echo"

	"github.com/ajinori/d4c/modules/text"
)

type Route struct {
	Path string
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error { return text.Module(c, "d4c") })

	e.Start(":8080")
}
