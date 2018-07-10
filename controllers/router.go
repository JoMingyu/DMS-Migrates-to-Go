package controllers

import (
	"DMS-Migrates-to-Go/controllers/account_student"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Setup 함수는 WAS의 비즈니스 로직 수행을 위한 API들을 모아 라우팅합니다.
func Setup(e *echo.Echo) {
	studentAPIGroup := e.Group("/student")
	adminAPIGroup := e.Group("/admin")

	tokenRequiredStudentAPIGroup := studentAPIGroup.Group("")
	tokenRequiredStudentAPIGroup.Use(middleware.JWT([]byte("secret")))

	tokenRequiredAdminAPIGroup := adminAPIGroup.Group("")
	tokenRequiredAdminAPIGroup.Use(middleware.JWT([]byte("secret")))

	studentAPIGroup.GET("/student/verify/id/:id", account_student.StudentCheckIDDuplication)
	studentAPIGroup.GET("/student/verify/uuid/:uuid", account_student.StudentValidateUUID)
	studentAPIGroup.GET("/student/verify/id/:id", account_student.StudentCheckIDDuplication)
	studentAPIGroup.POST("/student/signup", account_student.StudentSignup)
	studentAPIGroup.POST("/student/login", account_student.StudentLogin)
	tokenRequiredStudentAPIGroup.POST("/student/change-pw", account_student.ChangeStudentPassword)
}
