package config

type commonSetting struct {
	ServiceName string
}

type runSetting struct {
	Port string
}

var (
	CommonSetting = new(commonSetting)
	RunSetting    = new(runSetting)
	MongoURI      string
)

func init() {
	if CommonSetting == nil {
		CommonSetting.ServiceName = "dms-v2"
	}

	if RunSetting == nil {
		RunSetting.Port = "80"
	}

	if MongoURI == "" {
		MongoURI = "mongodb://localhost"
	}
}
