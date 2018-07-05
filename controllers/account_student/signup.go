package account_student

import (
	"DMS-Migrates-to-Go/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2-unstable/bson"
)

type studentSignupAPIBinder struct {
	Uuid string `json:"uuid"`
	Id   string `json:"id"`
	Pw   string `json:"pw"`
}

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

	// 학생 UUID 유효성 검사
	router.Add("GET", "/student/verify/uuid/:uuid", func(c echo.Context) error {
		uuid := c.Param("uuid")

		if count, _ := model.SignupWaitingCol.Find(bson.M{"uuid": uuid}).Count(); count != 0 {
			// UUID가 존재하는 경우
			return c.NoContent(http.StatusOK)
		}

		return c.NoContent(http.StatusNoContent)
	})

	// 학생 회원가입
	router.Add("POST", "/student/signup", func(c echo.Context) error {
		payload := &studentSignupAPIBinder{}

		if error := c.Bind(payload); error != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		var (
			uuid = payload.Uuid
			id   = payload.Id
			pw   = payload.Pw
		)

		if count, _ := model.StudentAccountCol.Find(bson.M{"id": id}).Count(); count != 0 {
			// ID가 이미 존재하는 경우
			return c.NoContent(http.StatusConflict)
		}

		signupWaiting := &model.SignupWaitingModel{}

		if error := model.SignupWaitingCol.Find(bson.M{"uuid": uuid}).One(signupWaiting); error != nil {
			// UUID가 존재하지 않는 경우
			return c.NoContent(http.StatusNoContent)
		}

		model.StudentAccountCol.Insert(model.StudentModel{
			Id:                    id,
			Pw:                    pw,
			Name:                  signupWaiting.Name,
			Number:                signupWaiting.Number,
			GoodPoint:             0,
			BadPoint:              0,
			PenaltyTrainingStatus: false,
			PenaltyLevel:          0,
		})

		model.SignupWaitingCol.Remove(bson.M{"uuid": payload.Uuid})

		return c.NoContent(http.StatusCreated)
	})
}
