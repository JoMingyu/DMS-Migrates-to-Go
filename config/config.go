package config

type commonSetting struct {
	ServiceName string
}

type runSetting struct {
	Port string
}

var (
	CommonSetting *commonSetting
	RunSetting    *runSetting
	MongoURI      string
)

func init() {
	if CommonSetting == nil {
		CommonSetting = &commonSetting{"dms-v2"}
	}

	if RunSetting == nil {
		RunSetting = &runSetting{"8000"}
	}

	if MongoURI == "" {
		MongoURI = "mongodb://localhost"
	}
}
