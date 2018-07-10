package main

import (
	"DMS-Migrates-to-Go/config"
	"DMS-Migrates-to-Go/controllers"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	controllers.Setup(e)

	e.Start(":" + strconv.Itoa(config.Port))
}
