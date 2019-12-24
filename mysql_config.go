package jotnar

// key - > main, slave
type baseMysql struct {
	Dsn     string
	MaxIdle int
	MaxOpen int
}

var MysqlConfig = make(map[string]*baseMysql)

func (j *Jotnar) InitMysql() *Jotnar {
	switch CurrentConfigType {
	case "default":
		readMysqlFromDefault()
	case "viper":
		readMysqlFromViper()
	}

	readMysqlFromViper()

	initGorm()

	return j
}

// support only one mysql instance, suggest use viper
func readMysqlFromDefault() {

}

// read mysql.main.xxx
// read mysql.slave.xxx
func readMysqlFromViper() {
	v := GetViper()

	if emptyStr(v.GetString("mysql.main.dsn")) {
		panic("must have mysql.main.dsn")
	}

	MysqlConfig["main"] = &baseMysql{
		Dsn:     v.GetString("mysql.main.dsn"),
		MaxIdle: v.GetInt("mysql.main.maxIdle"),
		MaxOpen: v.GetInt("mysql.main.maxOpen"),
	}

	if !emptyStr(v.GetString("mysql.slave.dsn")) {
		MysqlConfig["salve"] = &baseMysql{
			Dsn:     v.GetString("mysql.salve.dsn"),
			MaxIdle: v.GetInt("mysql.salve.maxIdle"),
			MaxOpen: v.GetInt("mysql.salve.maxOpen"),
		}
	}
}
