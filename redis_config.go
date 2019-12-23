package jotnar

type baseRedis struct {
	Network               string
	Address               string
	DialConnectionTimeout int
	DialReadTimeout       int
	DialWriteTimeout      int
	DialPassword          string
	DB                    int
	MaxIdle               int
	MaxActive             int
	TestOnBorrow          bool
	IdleTimeout           int
	Wait                  bool
}

var RedisConfig = make(map[string]*baseRedis)

func (j *Jotnar) InitRedisConfig() *Jotnar {
	
	return j
}