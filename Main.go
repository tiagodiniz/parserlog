package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// new echo instance
	e := echo.New()

	// set debug mode
	e.Debug = true

	// set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// add static paths
	e.Static("/", "static")


	// start server
	e.Start(":8888")

}
