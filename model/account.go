package model

import (
	"DMS-Migrates-to-Go/db"
)

type SignupWaitingModel struct {
	Uuid   string `bson:"uuid"`
	Name   string `bson:"name"`
	Number int    `bson:"number"`
}

var SignupWaitingCol = db.DB().C("signup_waiting")

type StudentModel struct {
	Id                    string `bson:"id"`
	Pw                    string `bson:"pw"`
	Name                  string `bson:"name"`
	Number                int    `bson:"number"`
	GoodPoint             int    `bson:"goodPoint"`
	BadPoint              int    `bson:"badPoint"`
	PenaltyTrainingStatus bool   `bson:"penaltyTrainingStatus`
	PenaltyLevel          int    `bson:"penaltyLevel"`
}

var StudentAccountCol = db.DB().C("account_student")

type Key struct {
	Owner     StudentModel `bson:"owner"`
	UserAgent string       `bson:"userAgent"`
}

type TokenModel struct {
	key      Key    `bson:"key"`
	identity string `bson:"identity"`
}

type AccessTokenModel TokenModel
type RefreshTokenModel TokenModel

var (
	AccessTokenCol  = db.DB().C("access_token")
	RefreshTokenCol = db.DB().C("refresh_token")
)
