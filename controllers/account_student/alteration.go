package account_student

import (
	"DMS-Migrates-to-Go/model"
	"DMS-Migrates-to-Go/util"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2-unstable/bson"
)

// ChangeStudentPassword 는 request body로 전달된 새로운 비밀번호로 학생 계정의 비밀번호를 변경합니다.
func ChangeStudentPassword(c echo.Context) error {
	type binder struct {
		CurrentPw string `json:"currentPw"`
		NewPw     string `json:"newPw"`
	}

	payload := &binder{}

	if err := c.Bind(payload); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if payload.CurrentPw == "" || payload.NewPw == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	student := util.ExtractStudentFromEchoContext(c)

	if student.Pw != payload.CurrentPw {
		return c.NoContent(http.StatusForbidden)
	}

	if payload.CurrentPw == payload.NewPw {
		// 현재 비밀번호와 새 비밀번호가 동일한 경우
		return c.NoContent(http.StatusConflict)
	}

	model.StudentAccountCol.Update(bson.M{"_id": student.ID}, bson.M{"$set": bson.M{"pw": payload.NewPw}})

	return c.NoContent(http.StatusOK)
}
