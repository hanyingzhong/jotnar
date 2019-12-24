# jotnar
a fast build framework about go application<br>
![](https://github.com/paulyung541/jotnar/workflows/.github/workflows/go.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

![](https://github.com/paulyung541/jotnar/blob/master/gopher_head.png)

## install
```sh
go get -u github.com/paulyung541/jotnar
```

* [config initialization and get value](https://github.com/paulyung541/jotnar#config-initialization-and-get-value)
  * [use default](https://github.com/paulyung541/jotnar#use-default)
  * [use viper](https://github.com/paulyung541/jotnar#use-viper)
* [mysql config](https://github.com/paulyung541/jotnar#mysql-config)
* [logger config](https://github.com/paulyung541/jotnar#logger-config)
* [redis config](https://github.com/paulyung541/jotnar#redis-config)

## config initialization and get value
there is 2 choice to use<br>
* the sdk default args packages
* use [viper](https://github.com/spf13/viper)

### use default
this way will launch your application use command line flag
run your application by this
```sh
./application --ip 192.168.0.1 --port 100
```

```go
jotnar.New().InitConfigDefaultCommandFlag()
fmt.Println(jotnar.GetVaule("ip"))
```

### use viper
must use `-f` flag
```sh
./application -f config.toml
```

your toml file like this
```toml
[server]
    url = "myserver.com"
```

code like this
```go
jotnar.New().InitConfigViperToml()
fmt.Println(jotnar.GetViper().GetString("server.url"))
```

for unit test
```go
func TestRead(t *testing.T) {
	jotnar.New().InitConfigViperTomlTest("config.toml")
	t.Log(jotnar.GetViper().GetString("server.url"))
}
```

## mysql config
write the config like this
```toml
[mysql.main]
    dsn = "root:@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
    maxIdle = 5
    maxOpen = 10
```

and the the init code like this, it's very easy to use
```go
package main

import (
	"fmt"

	"github.com/paulyung541/jotnar"
)

type Student struct {
	ID   uint32
	Name string
	Sex  uint8
}

func (*Student) TableName() string {
	return "student"
}

func main() {
	fmt.Println("this is example")

	jotnar.New().
		InitConfigViperToml().
		InitMysql().
		Init(func() {
			fmt.Println("initialization over")
		})

	var stu Student
	if err := jotnar.ReadGorm().First(&stu).Error; err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("stu = %+v\n", stu)
}
```

## logger config
log framework use [logrus](https://github.com/sirupsen/logrus), you just need add `InitLogger()` to initialization, and add the config in the `toml` file like follow
```toml
[log.default]
    file = ""
    level = "debug"
    format = "text"
    timeFormat = "2006-01-02 15:04:05.000000"
    isPretty = true
```

easy to use
```go
jotnar.GetLogger().Info("hello")
```

of course, you don't need any config file, there also can be run, just output to STDOUT

## redis config
use the [Redis client for Golang](https://github.com/go-redis/redis), toml file like follow
```toml
[redis.main]
    addr = "localhost:6379"
    password = ""
    db = 1
    maxRetries = 3
```

read redis use `ReadRedis()` and write redis use `WriteRedis()`

## License
[MIT](https://github.com/paulyung541/jotnar/blob/master/LICENSE)