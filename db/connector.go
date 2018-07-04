package db

import (
	"DMS-Migrates-to-Go/config"

	mgo "gopkg.in/mgo.v2-unstable"
)

var (
	session           *mgo.Session
	SignupWaitingCol  *mgo.Collection
	StudentAccountCol *mgo.Collection
	AccessTokenCol    *mgo.Collection
	RefreshTokenCol   *mgo.Collection
)

func init() {
	if session == nil {
		session, _ = mgo.Dial(config.MongoURI)
		session.SetSafe(new(mgo.Safe))
	}

	SignupWaitingCol = DB().C("signup_waiting")
	StudentAccountCol = DB().C("account_student")
	AccessTokenCol = DB().C("access_token")
	RefreshTokenCol = DB().C("refresh_token")
}

func DB() *mgo.Database {
	return session.DB(config.CommonSetting.ServiceName)
}
