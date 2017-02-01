package main

import (
	"fmt"

	"github.com/labstack/echo"
)

const port = ":3000"

func main() {
	e := echo.New()

	e.Static("/", "")
	e.Static("/", "client/dev")

	e.GET("/api", func(c echo.Context) error {
		return c.String(200, "1")
	})

	fmt.Println(port)

	if err := e.Start(port); err != nil {
		panic(err)
	}
}
