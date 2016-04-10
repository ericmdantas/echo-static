package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
)

func main() {
	e := echo.New()

	e.Get("/api", func(c echo.Context) error {
		return c.String(200, "1")
	})

	e.Get("/", echo.Static("client/dev"))
	e.Get("/*", echo.Static(""))

	e.Run(fasthttp.New(":3000"))
}
