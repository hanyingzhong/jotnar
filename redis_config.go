package jotnar

type baseRedis struct {
	Addr       string
	Password   string
	DB         int
	MaxRetries int
}

var RedisConfig = make(map[string]*baseRedis)

func (j *Jotnar) InitRedis() *Jotnar {
	switch CurrentConfigType {
	case "default":
		readRedisFromDefault()
	case "viper":
		readRedisFromViper()
	}

	initRedis()

	return j
}

func readRedisFromDefault() {

}

// read redis.main.xxx
// read redis.slave.xxx
func readRedisFromViper() {
	v := GetViper()

	if emptyStr(v.GetString("redis.main.addr")) {
		panic("must have redis.main.addr")
	}

	RedisConfig["main"] = &baseRedis{
		Addr:       v.GetString("redis.main.addr"),
		Password:   v.GetString("redis.main.password"),
		DB:         v.GetInt("redis.main.db"),
		MaxRetries: v.GetInt("redis.main.maxRetries"),
	}

	if !emptyStr(v.GetString("redis.slave.addr")) {
		RedisConfig["salve"] = &baseRedis{
			Addr:       v.GetString("redis.salve.addr"),
			Password:   v.GetString("redis.salve.password"),
			DB:         v.GetInt("redis.salve.db"),
			MaxRetries: v.GetInt("redis.salve.maxRetries"),
		}
	}
}
