package jotnar

import (
	"fmt"
	"os"
	"strings"
)

type config interface {
	witchConfig()
}

// if isCommandConfig true, then your application should
// run as follow
// ./application --key1 value1 \
// --key2 value2 \
// --key3 value3
// if isCommandConfig false, use config file to init like follow
// ./application -f configfile
type defaultConfig struct {
	IsCommandConfig bool
	Config          map[string]string // config key value
}

type viper struct {
	// viper.Viper
}

func (*defaultConfig) witchConfig() {}
func (*viper) witchConfig()         {}

func NewDefaultConfig() *defaultConfig {
	return &defaultConfig{
		IsCommandConfig: true,
		Config:          make(map[string]string),
	}
}

func NewViperConfig() *viper {
	return nil
}

var (
	DefaultConfig = NewDefaultConfig()
	ViperConfig   = NewViperConfig()
)

// after init, the config will load in memory
// then you guys can use by GetString...
func (j Jotnar) InitConfig(c config) Jotnar {
	switch cf := c.(type) {
	case *defaultConfig:
		doDefaultConfig(cf)
	case *viper:

	default:
		panic("should use defaultConfig or viper to init")
	}

	return j
}

func doDefaultConfig(cf *defaultConfig) {
	if cf.IsCommandConfig == true {
		if len(os.Args) < 2 {
			panic("there is no arg, please input args or do not use this.")
		}
		args := os.Args[1:]

		if len(args)%2 != 0 {
			panic("some arg dont have value, please input it.")
		}

		for i := 0; i < len(args); i += 2 {
			cf.Config[strings.TrimLeft(args[i], "-")] = args[i+1]
		}

		fmt.Printf("args = %+v\n", cf.Config)
	} else {
		// todo:
	}
}

// when use the DefaultConfig, you can use this to get value
// if dont have the key, then return <nil>
func GetValue(key string) string {
	if v, ok := DefaultConfig.Config[key]; !ok {
		return "<nil>"
	} else {
		return v
	}
}
