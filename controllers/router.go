package controllers

import (
	"DMS-Migrates-to-Go/controllers/account_student"

	"github.com/labstack/echo"
)

func Setup(router *echo.Router) {
	account_student.Setup(router)
}
