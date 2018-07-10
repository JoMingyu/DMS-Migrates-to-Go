package main

import (
	"DMS-Migrates-to-Go/config"
	"DMS-Migrates-to-Go/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	controllers.Setup(e.Router(), e.Group(""))

	e.Start(":" + strconv.Itoa(config.Port))
}
