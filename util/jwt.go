package util

import (
	"DMS-Migrates-to-Go/model"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	uuid "github.com/nu7hatch/gouuid"
	"gopkg.in/mgo.v2-unstable/bson"
)

func generateUUID() string {
	identity, _ := uuid.NewV4()
	identityStr := string(identity[:])

	return identityStr
}

func GenerateAccessToken(owner model.StudentModel, userAgent string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	identity := generateUUID()

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, _ := token.SignedString([]byte("secret"))

	model.AccessTokenCol.Insert(model.AccessTokenModel{
		Key: model.Key{
			Owner:     owner,
			UserAgent: userAgent,
		},
		Identity: identity,
	})

	return t
}

func GenerateRefreshToken(owner model.StudentModel, userAgent string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	identity := generateUUID()

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	t, _ := token.SignedString([]byte("secret"))

	model.RefreshTokenCol.Insert(model.RefreshTokenModel{
		Key: model.Key{
			Owner:     owner,
			UserAgent: userAgent,
		},
		Identity: identity,
	})

	return t
}

func ExtractStudentFromEchoContext(c echo.Context) *model.StudentModel {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	identity := claims["identity"]
	token := &model.AccessTokenModel{}

	if err := model.AccessTokenCol.Find(bson.M{"identity": identity}).One(token); err != nil {
		return nil
	}

	return &token.Key.Owner
}
