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
		ID string `json:"id"`
		Pw string `json:"pw"`
	}

	payload := &binder{}

	if err := c.Bind(payload); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if payload.ID == "" || payload.Pw == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	student := &model.StudentModel{}
	if err := model.StudentAccountCol.Find(bson.M{"_id": payload.ID, "pw": payload.Pw}).One(student); err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken":  util.GenerateAccessToken(student, ""),
		"refreshToken": util.GenerateRefreshToken(student, ""),
	})
}
