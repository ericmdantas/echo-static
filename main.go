package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const port = ":3000"

func main() {
	e := echo.New()

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "client/dev",
	}))

	e.Use(middleware.Static(""))
	e.Use(middleware.Static("client/dev"))

	e.GET("/api", func(c echo.Context) error {
		return c.String(200, "1")
	})

	fmt.Println(port)

	if err := e.Start(port); err != nil {
		panic(err)
	}
}
