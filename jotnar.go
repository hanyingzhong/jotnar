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
