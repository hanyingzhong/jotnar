package jotnar

type logConfig struct {
	FilePath   string
	Level      string
	Format     string // text or json
	Timeformat string
	IsPretty   bool
}

var defualtLogConfig *logConfig

func (j *Jotnar) InitLogger() *Jotnar {
	switch CurrentConfigType {
	case "default":
		logReadFromDefault()
	case "viper":
		logReadFromViper()
	}

	initLogrus()

	return j
}

func logReadFromDefault() {
	defualtLogConfig = &logConfig{
		Level:      "debug",
		Format:     "text",
		Timeformat: "2006-01-02 15:04:05",
		IsPretty:   true,
	}
}

func logReadFromViper() {
	v := GetViper()
	defualtLogConfig = &logConfig{
		FilePath:   v.GetString("log.default.file"),
		Level:      GetString("log.default.level", "info"),
		Format:     GetString("log.default.format", "text"),
		Timeformat: GetString("log.default.timeFormat", "2006-01-02 15:04:05"),
		IsPretty:   v.GetBool("log.default.isPretty"),
	}
}
