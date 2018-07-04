package config

type commonSetting struct {
	ServiceName string
}

type runSetting struct {
	Port string
}

var (
	CommonSetting = commonSetting{"dms-v2"}
	RunSetting    = runSetting{"4000"}
	MongoURI      = "mongodb://localhost"
)
