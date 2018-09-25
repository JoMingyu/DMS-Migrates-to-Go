package model

import (
	"DMS-Migrates-to-Go/db"

	"gopkg.in/mgo.v2-unstable/bson"
)

type PostModel struct {
	ID      bson.ObjectId `bson:"_id"`
	Author  string
	Title   string
	Content string
	Pinned  string
}

var (
	FAQCol    = db.DB().C("post_faq")
	NoticeCol = db.DB().C("post_notice")
	RuleCol   = db.DB().C("post_rule")
)
