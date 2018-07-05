package config

type commonSetting struct {
	ServiceName string
}

type runSetting struct {
	Port string
}

var (
	// CommonSetting 은 WAS에 필요한 기본적인 정보들을 관리합니다.
	CommonSetting *commonSetting

	// RunSetting 은 서버 run 시 필요한 정보들을 제공합니다.
	RunSetting *runSetting

	// MongoURI 는 mgo의 세션을 위한 MongoDB URI를 명시합니다.
	MongoURI string
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
