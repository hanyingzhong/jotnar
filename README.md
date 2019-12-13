# jotnar
a fast build framework about go application<br>
[![Build Status](https://travis-ci.org/jmoiron/sqlx.svg?branch=master)](https://travis-ci.org/jmoiron/sqlx)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)
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
```
func TestRead(t *testing.T) {
	jotnar.New().InitConfigViperTomlTest("config.toml")
	t.Log(jotnar.GetViper().GetString("server.url"))
}
```

## License
[MIT](https://github.com/paulyung541/jotnar/blob/master/LICENSE)