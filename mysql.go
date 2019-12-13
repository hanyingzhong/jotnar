package jotnar

type mysqlConfig struct {
	mysqlSet map[string]*baseMysql
}

type baseMysql struct {
	dsn     string
	maxIdle int
	maxOpen int
}

func (j *Jotnar) InitMysql() *Jotnar {
	return j
}
