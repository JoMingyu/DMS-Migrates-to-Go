package account_student

import (
	"DMS-Migrates-to-Go/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2-unstable/bson"
)

func Setup(router *echo.Router) {
	// 학생 ID 중복체크
	router.Add("GET", "/student/verify/id/:id", func(c echo.Context) error {
		id := c.Param("id")

		if count, _ := model.StudentAccountCol.Find(bson.M{"id": id}).Count(); count == 0 {
			// ID가 존재하지 않는 경우
			return c.NoContent(http.StatusOK)
		}

		return c.NoContent(http.StatusConflict)
	})

	router.Add("GET", "/student/verify/uuid/:uuid", func(c echo.Context) error {
		uuid := c.Param("uuid")

		if count, _ := db.SignupWaitingCol.Find(bson.M{"uuid": uuid}).Count(); count != 0 {
			return c.String(http.StatusConflict, "")
		} else {
			return c.String(http.StatusOK, "")
		}
	})
}
