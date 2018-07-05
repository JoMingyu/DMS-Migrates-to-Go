package account_student

import (
	"DMS-Migrates-to-Go/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2-unstable/bson"
)

// StudentLogin provides student login
func StudentLogin(c echo.Context) error {
	payload := &studentLoginAPIBinder{}

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
