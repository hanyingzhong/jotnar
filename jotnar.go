// this repo just direct depend on the follow 3 frameworks
//
// github.com/spf13/viper
// github.com/sirupsen/logrus
// github.com/jinzhu/gorm
//
// the other module like redis mongo client can use InitRedis
// or InitMongoDB, and choice the client you like in your project
// use the jotnar config struct to Init your client.
package jotnar

import (
	"fmt"
	"os"
	"runtime"
)

type Jotnar struct{}

func New() *Jotnar {
	defaultInit()
	return new(Jotnar)
}

func defaultInit() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// customer common initialization
func (*Jotnar) Init(functions ...func()) {
	for _, f := range functions {
		if f != nil {
			f()
		}
	}
}

func emptyStr(s string) bool {
	return s == ""
}

func errExit(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("program exit...")
		os.Exit(-1)
	}
}
