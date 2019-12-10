# jotnar
a fast build framework about go application 

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

### default init
```go
jotnar.New().InitConfig(jotnar.DefaultConfig)
fmt.Println(jotnar.GetVaule("ip"))
```
run your application by this
```sh
./application --ip 192.168.0.1 --port 100
```

* [MIT License](https://github.com/paulyung541/jotnar/blob/master/LICENSE)