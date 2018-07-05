package controllers

import (
	"DMS-Migrates-to-Go/controllers/account_student"

	"github.com/labstack/echo"
)

// Setup 함수는 WAS의 비즈니스 로직 수행을 위한 API들을 모아 라우팅합니다.
func Setup(router *echo.Router) {
	router.Add("GET", "/student/verify/id/:id", account_student.StudentCheckIDDuplication)
	router.Add("GET", "/student/verify/uuid/:uuid", account_student.StudentValidateUUID)
	router.Add("GET", "/student/verify/id/:id", account_student.StudentSignup)
	router.Add("GET", "/student/login", account_student.StudentLogin)
}
