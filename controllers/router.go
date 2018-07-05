package controllers

import (
	"DMS-Migrates-to-Go/controllers/account_student"

	"github.com/labstack/echo"
)

// Setup routes all APIs
func Setup(router *echo.Router) {
	router.Add("GET", "/student/verify/id/:id", account_student.StudentCheckIDDuplication)
	router.Add("GET", "/student/verify/uuid/:uuid", account_student.StudentValidateUUID)
	router.Add("GET", "/student/verify/id/:id", account_student.StudentSignup)
	router.Add("GET", "/student/login", account_student.StudentLogin)
}
