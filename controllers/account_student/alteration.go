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
		CurrentPW string `json:"currentPw"`
		NewPW     string `json:"newPw"`
	}

	payload := &binder{}

	if err := c.Bind(payload); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	var (
		currentPw = payload.CurrentPW
		newPw     = payload.NewPW
	)

	if currentPw == "" || newPw == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	student := util.ExtractStudentFromEchoContext(c)

	if student.Pw != currentPw {
		return c.NoContent(http.StatusForbidden)
	}

	if currentPw == newPw {
		// 현재 비밀번호와 새 비밀번호가 동일한 경우
		return c.NoContent(http.StatusConflict)
	}

	model.StudentAccountCol.Update(bson.M{"id": student.Id}, bson.M{"$set": bson.M{"pw": newPw}})

	return c.NoContent(http.StatusOK)
}
