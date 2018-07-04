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

	controllers.Setup(e.Router())

	err := e.Start(":" + config.RunSetting.Port)
	if err != nil {
		panic(err)
	}
}
