package main

import (
	"fmt"
	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/fasthttp"
)

const port = ":3000"

func main() {
	e := echo.New()

	e.Use(m.Static(""))
	e.Use(m.Static("client/dev"))	

	e.Get("/api", func(c echo.Context) error {
		return c.String(200, "1")
	})
	
	fmt.Println(port)

	e.Run(fasthttp.New(port))
}
