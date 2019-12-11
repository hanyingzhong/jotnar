package jotnar

type mysqlConfig struct {
	mysqlSet map[string]*baseMysql
}

type baseMysql struct {
	dsn string
}
