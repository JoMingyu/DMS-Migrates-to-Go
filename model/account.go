package model

import (
	"DMS-Migrates-to-Go/db"

	"gopkg.in/mgo.v2-unstable/bson"
)

// SignupWaitingModel 구조체는 아직 회원가입되지 않은 학생 데이터를 관리합니다.
type SignupWaitingModel struct {
	Uuid   string `bson:"_id"` // PK
	Name   string
	Number int
}

// StudentModel 구조체는 학생 계정의 데이터를 관리합니다.
type StudentModel struct {
	ID                    string `bson:"_id"` // PK
	Pw                    string
	Name                  string
	Number                int
	GoodPoint             int
	BadPoint              int
	PenaltyTrainingStatus bool
	PenaltyLevel          int
}

// Key 는 JWT token이 각 계정의 여러 디바이스에 하나씩 할당되도록 하기 위해 TokenModel에서 사용할 구조체입니다.
type Key struct {
	Owner     StudentModel
	UserAgent string
}

// TokenModel 구조체는 JWT access token과 refresh token에 대한 데이터를 관리합니다.
type TokenModel struct {
	ID  bson.ObjectId `bson:"_id"`
	Key Key           // Unique
}

var (
	// SignupWaitingCol 은 SignupWaitingModel에 대한 collection을 참조합니다.
	SignupWaitingCol = db.DB().C("signup_waiting")

	// StudentAccountCol 은 StudentAccountModel에 대한 collection을 침조합니다.
	StudentAccountCol = db.DB().C("account_student")

	// AccessTokenCol 은 AccessTokenModel에 대한 collection을 참조합니다.
	AccessTokenCol = db.DB().C("access_token")

	// RefreshTokenCol 은 RefreshTokenModel에 대한 collection을 참조합니다.
	RefreshTokenCol = db.DB().C("refresh_token")
)
