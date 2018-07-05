package db

import (
	"DMS-Migrates-to-Go/config"

	mgo "gopkg.in/mgo.v2-unstable"
)

var (
	session *mgo.Session
)

func init() {
	if session == nil {
		session, _ = mgo.Dial(config.MongoURI)
		session.SetSafe(new(mgo.Safe))
	}
}

func DB() *mgo.Database {
	return session.DB(config.CommonSetting.ServiceName)
}
