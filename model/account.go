package model

type SignupWaitingModel struct {
	Uuid   string `bson:"uuid"`
	Name   string `bson:"name"`
	Number int    `bson:"number"`
}

type StudentModel struct {
	Id                    string `bson:"id"`
	Pw                    string `bson:"pw"`
	Name                  string `bson:"name"`
	GoodPoint             int    `bson:"goodPoint"`
	BadPoint              int    `bson:"badPoint"`
	PenaltyTrainingStatus bool   `bson:"penaltyTrainingStatus`
	PenaltyLevel          int    `bson:"penaltyLevel"`
}

type Key struct {
	Owner     StudentModel `bson:"owner"`
	UserAgent string       `bson:"userAgent"`
}

type TokenModel struct {
	key      Key    `bson:"key"`
	identity string `bson:"identity"`
}

type AccessTokenModel struct {
	key      Key    `bson:"key"`
	identity string `bson:"identity"`
}

type RefreshTokenModel struct {
	key      Key    `bson:"key"`
	identity string `bson:"identity"`
}
