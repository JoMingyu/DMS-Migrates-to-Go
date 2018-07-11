package account_student

import (
	"DMS-Migrates-to-Go/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2-unstable/bson"
)

// StudentCheckIDDuplication 함수는 학생 계정에 대해 ID의 중복 여부를 체크합니다
func StudentCheckIDDuplication(c echo.Context) error {
	id := c.Param("id")

	if count, _ := model.StudentAccountCol.Find(bson.M{"_id": id}).Count(); count == 0 {
		// ID가 존재하지 않는 경우
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusConflict)
}

// StudentValidateUUID 함수는 UUID의 유효성(존재 여부)을 체크합니다.
func StudentValidateUUID(c echo.Context) error {
	uuid := c.Param("uuid")

	if count, _ := model.SignupWaitingCol.Find(bson.M{"_id": uuid}).Count(); count != 0 {
		// UUID가 존재하는 경우
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusNoContent)
}

// StudentSignup 함수는 학생 계정 회원가입을 수행합니다.
func StudentSignup(c echo.Context) error {
	type binder struct {
		Uuid string `json:"uuid"`
		ID   string `json:"id"`
		Pw   string `json:"pw"`
	}

	payload := &binder{}

	if err := c.Bind(payload); err != nil {
		return err
	}

	if payload.Uuid == "" || payload.ID == "" || payload.Pw == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	if count, _ := model.StudentAccountCol.Find(bson.M{"_id": payload.ID}).Count(); count != 0 {
		// ID가 이미 존재하는 경우
		return c.NoContent(http.StatusConflict)
	}

	signupWaiting := &model.SignupWaitingModel{}
	if err := model.SignupWaitingCol.Find(bson.M{"_id": payload.Uuid}).One(signupWaiting); err != nil {
		// UUID가 존재하지 않는 경우
		return c.NoContent(http.StatusNoContent)
	}

	model.StudentAccountCol.Insert(model.StudentModel{
		ID:                    payload.ID,
		Pw:                    payload.Pw,
		Name:                  signupWaiting.Name,
		Number:                signupWaiting.Number,
		GoodPoint:             0,
		BadPoint:              0,
		PenaltyTrainingStatus: false,
		PenaltyLevel:          0,
	})

	model.SignupWaitingCol.Remove(bson.M{"_id": payload.Uuid})

	return c.NoContent(http.StatusCreated)
}
