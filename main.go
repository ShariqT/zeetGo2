package main

import (
	"github.com/labstack/echo"
)
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "test") 
	})
	e.Start(":3000")
}
