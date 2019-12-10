# jotnar
a fast build framework about go application
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)

![](https://github.com/paulyung541/jotnar/blob/master/gopher_head.png)

## install
```sh
go get -u github.com/paulyung541/jotnar
```

* [config initialization and get value]()

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
jotnar.New().InitConfig(jotnar.DefaultConfig)
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
jotnar.New().InitConfig(jotnar.ViperConfigToml)
fmt.Println(jotnar.GetViper().GetString("server.url"))
```

## License
[MIT](https://github.com/paulyung541/jotnar/blob/master/LICENSE)