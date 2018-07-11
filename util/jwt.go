package util

import (
	"DMS-Migrates-to-Go/model"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2-unstable"
	"gopkg.in/mgo.v2-unstable/bson"
)

func generateToken(owner *model.StudentModel, userAgent string, collection *mgo.Collection, expire time.Duration) string {
	token := jwt.New(jwt.SigningMethodHS256)
	identity := bson.NewObjectId()

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity.String()
	claims["exp"] = time.Now().Add(expire).Unix()

	t, _ := token.SignedString([]byte("secret"))

	key := model.Key{
		Owner:     *owner,
		UserAgent: userAgent,
	}

	collection.RemoveAll(bson.M{"key": key})

	collection.Insert(model.TokenModel{
		ID:  identity,
		Key: key,
	})

	return t
}

func GenerateAccessToken(owner *model.StudentModel, userAgent string) string {
	return generateToken(owner, userAgent, model.AccessTokenCol, time.Hour*24)
}

func GenerateRefreshToken(owner *model.StudentModel, userAgent string) string {
	return generateToken(owner, userAgent, model.RefreshTokenCol, time.Hour*24*30)
}

func ExtractStudentFromEchoContext(c echo.Context) *model.StudentModel {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	identity := claims["identity"]
	token := &model.TokenModel{}

	if err := model.AccessTokenCol.Find(bson.M{"_id": bson.ObjectIdHex(identity.(string))}).One(token); err != nil {
		return nil
	}

	return &token.Key.Owner
}
