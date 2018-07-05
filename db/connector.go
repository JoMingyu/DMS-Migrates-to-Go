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

// DB 함수는 db.session에 대해 데이터베이스의 참조를 얻어 반환합니다.
func DB() *mgo.Database {
	return session.DB(config.CommonSetting.ServiceName)
}
