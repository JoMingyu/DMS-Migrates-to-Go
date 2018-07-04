package db

import (
	"DMS-Migrates-to-Go/config"

	mgo "gopkg.in/mgo.v2-unstable"
)

var Session *mgo.Session

func init() {
	if Session == nil {
		Session, _ = mgo.Dial(config.MongoURI)
		Session.SetSafe(new(mgo.Safe))
	}
}

func DB() *mgo.Database {
	return Session.DB(config.CommonSetting.ServiceName)
}
