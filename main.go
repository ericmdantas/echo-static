package main

import (
	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/fasthttp"
)

func main() {
	e := echo.New()

	e.Use(m.Static(""))
	e.Use(m.Static("client/dev"))	

	e.Get("/api", func(c echo.Context) error {
		return c.String(200, "1")
	})

	e.Run(fasthttp.New(":3000"))
}
