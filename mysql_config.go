package jotnar

// key - > main, slave
type mysqlConfig struct {
	MysqlSet map[string]*baseMysql
}

type baseMysql struct {
	Dsn     string
	MaxIdle int
	MaxOpen int
}

var MConfig *mysqlConfig

func (j *Jotnar) InitMysql() *Jotnar {
	switch CurrentConfigType {
	case "default":
		readFromDefault()
	case "viper":
		readFromViper()
	}

	// init
	MConfig = readFromViper()

	// feture: support to choice orm framework
	InitGorm()

	return j
}

// support only one mysql instance, suggest use viper
func readFromDefault() {

}

// read all mysql.main.xxx
// read all mysql.slave.xxx
func readFromViper() *mysqlConfig {
	v := GetViper()

	if emptyStr(v.GetString("mysql.main.dsn")) {
		panic("must have mysql.main.dsn")
	}

	mc := &mysqlConfig{make(map[string]*baseMysql)}
	mc.MysqlSet["main"] = &baseMysql{
		Dsn:     v.GetString("mysql.main.dsn"),
		MaxIdle: v.GetInt("mysql.main.maxIdle"),
		MaxOpen: v.GetInt("mysql.main.maxOpen"),
	}

	if !emptyStr(v.GetString("mysql.slave.dsn")) {
		mc.MysqlSet["salve"] = &baseMysql{
			Dsn:     v.GetString("mysql.salve.dsn"),
			MaxIdle: v.GetInt("mysql.salve.maxIdle"),
			MaxOpen: v.GetInt("mysql.salve.maxOpen"),
		}
	}

	return mc
}
