package model

import (
	"DMS-Migrates-to-Go/db"

	"gopkg.in/mgo.v2-unstable/bson"
)

type PointRuleModel struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	PointType bool
	MinPoint  int
	MaxPoint  int
}

type PointHistoryModel struct {
	ID        bson.ObjectId `bson:"_id"`
	Student   StudentModel
	Reason    string
	PointType bool
	Point     int
}

var (
	PointRuleCol    = db.DB().C("point_rule")
	PointHistoryCol = db.DB().C("point_history")
)
