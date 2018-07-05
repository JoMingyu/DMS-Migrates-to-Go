package account_student

import (
	"DMS-Migrates-to-Go/model"
	"DMS-Migrates-to-Go/util"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2-unstable/bson"
)

// StudentLogin 함수는 학생 계정 로그인을 수행합니다.
func StudentLogin(c echo.Context) error {
	type binder struct {
		Id string `json:"id"`
		Pw string `json:"pw"`
	}

	payload := &binder{}

	if err := c.Bind(payload); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	var (
		id = payload.Id
		pw = payload.Pw
	)

	if id == "" || pw == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	student := &model.StudentModel{}
	if err := model.StudentAccountCol.Find(bson.M{"id": id, "pw": pw}).One(student); err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.NoContent(http.StatusOK)
}
