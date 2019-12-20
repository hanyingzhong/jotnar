package jotnar

type logConfig struct {
	FilePath   string
	Level      string
	Format     string // text or json
	Timeformat string
}

var defualtLogConfig *logConfig

func InitLogger() {
	switch CurrentConfigType {
	case "default":
		logReadFromDefault()
	case "viper":
		logReadFromViper()
	}
}

func logReadFromDefault() {

}

func logReadFromViper() {
	v := GetViper()
	defualtLogConfig = &logConfig{
		FilePath: v.GetString("log.default.file"),
		Level: v.GetString("log.default.level"),
		Format: v.GetString("log.default.format"),
		Timeformat: v.GetString("log.default.timeFormat"),
	}
}
